package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/codeready-toolchain/operator-role-linter/download"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/discovery"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig *restclient.Config

func main() {
	var err error
	kubeconfig, err = contextConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var skippedGroups string
	var outputFile string

	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "download the api resources from the given openshift cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			// find the API for the given resource type
			client, err := discovery.NewDiscoveryClientForConfig(kubeconfig)
			if err != nil {
				return err
			}
			resources, err := download.Download(client, skippedGroups)
			if err != nil {
				return err
			}
			// now, dump the resources into a YAML file
			data, err := yaml.Marshal(resources)
			if err != nil {
				return err
			}
			outputDir := filepath.Dir(outputFile)
			err = os.Mkdir(outputDir, os.FileMode(0755))
			if err != nil && !os.IsExist(err) {
				return err
			}
			return ioutil.WriteFile(outputFile, data, os.FileMode(0644))
		},
		SilenceUsage: true,
	}
	flags := downloadCmd.Flags()
	flags.StringVar(&skippedGroups, "skippedGroups", "", "a comma-separated list of API groups to skip")
	flags.StringVar(&outputFile, "output-file", "", "the output dir in which API resources (JSON files) will be downloaded")
	downloadCmd.MarkFlagRequired("output-file")

	kubeConfigFlags := genericclioptions.NewConfigFlags(true)
	kubeConfigFlags.AddFlags(downloadCmd.PersistentFlags())
	if err := downloadCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

// contextConfig returns the config of the current context in kubeconfig
func contextConfig() (*restclient.Config, error) {
	var kubeconfig *string
	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	if home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	// flag.Parse()
	// use the current context in kubeconfig
	return clientcmd.BuildConfigFromFlags("", *kubeconfig)
}
