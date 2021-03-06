/*
Copyright © 2020 Appvia Ltd <info@appvia.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Allows you to get one of more resources in statuspage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("statuspage get error: missing required argument. See 'statuspage get -h' for help.")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
