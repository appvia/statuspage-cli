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
	"../utils"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var getPageCmd = &cobra.Command{
	Use:   "page",
	Short: "Get a list of pages.",
	Run: func(cmd *cobra.Command, args []string) {
		apiKeyFromEnv := utils.GetEnv("API_KEY")
		apiKey = utils.DefaultToEnv(apiKeyFromEnv, apiKey)

		client := &http.Client{Timeout: time.Second * 10}
		request, err := http.NewRequest("GET", "https://api.statuspage.io/v1/pages", nil)
		request.Header.Set("Authorization", "OAuth "+apiKey)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		resp, err := client.Do(request)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(string(body))
	},
}

func init() {
	getPageCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "API_KEY environment variable. API key to authenticate against the status page API (required)")
	getCmd.AddCommand(getPageCmd)
}
