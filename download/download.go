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

package download

import (
	"fmt"
	"strings"

	"github.com/codeready-toolchain/operator-role-linter/pkg/apiresources"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

// Download downloads the API Resources from the given cluster URL
func Download(client discovery.DiscoveryInterface, skippedGroups string) (apiresources.APIResources, error) {
	// list of resources by APIGroup
	resources := apiresources.APIResources{}

	excludedGroups := map[string]bool{}
	for _, g := range strings.Split(skippedGroups, ",") {
		excludedGroups[g] = true
	}

	apiGroups, apiResourceLists, err := client.ServerGroupsAndResources()
	if err != nil {
		return apiresources.APIResources{}, err
	}

	preferredGroupVersions := make(map[string]string)
	for _, apiGroup := range apiGroups {
		fmt.Printf("apiGroup: %s -> %s\n", apiGroup.Name, apiGroup.PreferredVersion.Version)
		preferredGroupVersions[apiGroup.Name] = apiGroup.PreferredVersion.Version
	}

	for _, apiResourceList := range apiResourceLists {
		gv, err := schema.ParseGroupVersion(apiResourceList.GroupVersion)
		if err != nil {
			return apiresources.APIResources{}, err
		}
		group := gv.Group
		if version, found := preferredGroupVersions[group]; !found || version != gv.Version {
			continue
		}
		// fmt.Printf("apiResourceList: APIVersion=%s / GroupVersion=%s\n", apiResourceList.APIVersion, apiResourceList.GroupVersion)
		fmt.Printf("group '%s': %s\n", group, gv.Version)
		// spew.Dump(apiResourceList)
		// check if group is skipped
		if _, excluded := excludedGroups[group]; group != "" && excluded {
			log.Infof("skipping group '%s'", group)
			continue
		}
		resources[group] = make(map[string]apiresources.Resource)
		for _, apiResource := range apiResourceList.APIResources {
			resources[group][apiResource.Name] = apiresources.Resource{
				Namespaced: apiResource.Namespaced,
				Verbs:      apiResource.Verbs,
			}
		}
	}

	return resources, nil
}
