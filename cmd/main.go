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

package main

import (
	"fmt"
	"os"

	"github.com/codeready-toolchain/operator-role-linter/pkg/apiresources"

	"github.com/spf13/cobra"
)

func main() {
	var roleFile string
	var crdsPath string
	var clusterVersion string
	verifyCmd := &cobra.Command{
		Use:   "operator-role-linter",
		Short: "verifies that the deploy/roles.yaml and deploy/cluster-roles.yaml are valid",
		Long: `Verifies that the roles in the deploy/roles.yaml are namespaced roles, and
	that roles in the deploy/cluster-roles.yaml are cluster roles.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			errors, err := apiresources.Verify(roleFile, crdsPath, clusterVersion)
			if err != nil {
				return err
			}
			for _, e := range errors {
				fmt.Println(e.Error())
			}
			return nil
		},
	}
	flags := verifyCmd.Flags()
	flags.StringVar(&roleFile, "role-file", "", "the path to the role/clusterrole to check")
	flags.StringVar(&crdsPath, "crds-dir", "", "the path to the directory with the custom CRDs to include")
	flags.StringVar(&clusterVersion, "cluster-version", "openshift-4.2", "the version of the target cluster (eg: 'openshift-4.2'")
	verifyCmd.MarkFlagRequired("role-file")
	verifyCmd.MarkFlagRequired("crds-dir")
	if err := verifyCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
