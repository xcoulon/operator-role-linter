// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/codeready-toolchain/operator-role-linter/pkg/apiresources"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the lintCmd.
func Execute() {
	var clusterURL string
	var token string
	var skippedGroups string
	var outputDir string

	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "download the api resources from the given openshift cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			config := tls.Config{
				InsecureSkipVerify: true,
			}
			tr := &http.Transport{TLSClientConfig: &config}
			client := &http.Client{Transport: tr}
			resources, err := Download(clusterURL, token, skippedGroups, client)
			if err != nil {
				return err
			}
			// now, dump the resources into a YAML file
			data, err := yaml.Marshal(resources)
			if err != nil {
				return err
			}
			return ioutil.WriteFile(filepath.Join(outputDir, "apiresources.yaml"), data, 0644)
		},
		SilenceUsage: true,
	}
	flags := downloadCmd.Flags()
	flags.StringVar(&clusterURL, "cluster-url", "", "the URL of the cluster to call to download the API resources")
	flags.StringVar(&token, "token", "", "the auth token (see 'oc whoami -t')")
	flags.StringVar(&skippedGroups, "skippedGroups", "", "a comma-separated list of API groups to skip")
	flags.StringVar(&outputDir, "output-dir", "o", "the output dir in which API resources (JSON files) will be downloaded")
	downloadCmd.MarkFlagRequired("cluster-url")
	downloadCmd.MarkFlagRequired("token")
	downloadCmd.MarkFlagRequired("output-dir")
	if err := downloadCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// APIResourceList an API Resource list for a given group
type APIResourceList struct {
	Resources []Resource
}

// Resource an API Group resource
type Resource struct {
	Name       string   `yaml:"name"`
	Namespaced bool     `yaml:"namespaced"`
	Verbs      []string `yaml:"verbs"`
}

// APIGroups the API groups
type APIGroups struct {
	Groups []APIGroup
}

// APIGroup an API group
type APIGroup struct {
	Name             string
	PreferredVersion PreferredVersion `yaml:"preferredVersion"`
}

// PreferredVersion the preferred version, as an API group may have multiple versions
type PreferredVersion struct {
	GroupVersion string `yaml:"groupVersion"`
	Version      string `yaml:"version"`
}

// Download downloads the API Resources from the given cluster URL
func Download(clusterURL, token, skippedGroups string, client *http.Client) (apiresources.APIResources, error) {
	// list of resources by APIGroup
	resources := apiresources.APIResources{}
	// process the `/api/v1` API group
	l := APIResourceList{}
	err := process(clusterURL+"/api/v1", token, client, &l)
	if err != nil {
		return apiresources.APIResources{}, errors.Wrapf(err, "unable to collect all resources")
	}
	resources[""] = make(map[string]apiresources.Resource, len(l.Resources))
	for _, r := range l.Resources {
		resources[""][r.Name] = apiresources.Resource{
			Namespaced: r.Namespaced,
			Verbs:      r.Verbs,
		}
	}
	// now, collect all other API groups and process each one of them
	groups := APIGroups{}
	err = process(clusterURL+"/apis", token, client, &groups)
	if err != nil {
		return apiresources.APIResources{}, errors.Wrapf(err, "unable to collect all resources")
	}
	excludedGroups := map[string]bool{}
	for _, g := range strings.Split(skippedGroups, ",") {
		excludedGroups[g] = true
	}
	for _, g := range groups.Groups {
		// check if group is skipped
		if _, excluded := excludedGroups[g.Name]; excluded {
			log.Infof("skipping group '%s'", g.Name)
			continue
		}
		l := APIResourceList{}
		err = process(clusterURL+"/apis/"+g.PreferredVersion.GroupVersion, token, client, &l)
		if err != nil {
			return apiresources.APIResources{}, errors.Wrapf(err, "unable to collect all resources")
		}
		resources[g.Name] = make(map[string]apiresources.Resource, len(l.Resources))
		for _, r := range l.Resources {
			resources[g.Name][r.Name] = apiresources.Resource{
				Namespaced: r.Namespaced,
				Verbs:      r.Verbs,
			}
		}
	}
	return resources, nil
}

func process(url, token string, client *http.Client, data interface{}) error {
	log.Info("collecting data from ", url)
	content, err := get(url, token, client)
	// log.Debug(string(content))
	if err != nil {
		return errors.Wrapf(err, "unable to process endpoint %s", url)
	}
	err = yaml.Unmarshal(content, data)
	if err != nil {
		return errors.Wrapf(err, "unable to process endpoint %s", url)
	}
	return nil
}

func get(url, token string, client *http.Client) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/yaml")
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to download %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, errors.Errorf("unable to download: %s", resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}
