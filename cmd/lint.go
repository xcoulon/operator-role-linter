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
	"fmt"
	"os"

	"github.com/codeready-toolchain/operator-role-linter/pkg/apiresources"

	"github.com/spf13/cobra"
)

var ruleFile string
var crdsPath string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the verifyCmd.
func Execute() {
	verifyCmd := &cobra.Command{
		Use:   "operator-role-linter",
		Short: "verifies that the deploy/roles.yaml and deploy/cluster-roles.yaml are valid",
		Long: `Verifies that the roles in the deploy/roles.yaml are namespaced roles, and
	that roles in the deploy/cluster-roles.yaml are cluster roles.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			errors, err := apiresources.Verify(ruleFile, crdsPath)
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
	flags.StringVar(&ruleFile, "rbac", "", "the path to the main RBAC file to check")
	flags.StringVar(&crdsPath, "crd", "", "the path to the directory with the custom CRDs to include")
	// TODO: include path to CRDs
	verifyCmd.MarkFlagRequired("path")
	if err := verifyCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
