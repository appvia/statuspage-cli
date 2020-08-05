/*
Copyright © 2020 Mohammud Yassine Jaffoo <mohammudyassine.jaffoo@appvia.io>

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
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"../utils"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var componentID string
var componentDescription string
var componentStatus string
var componentName string
var componentShowcase bool
var url string

var getComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Get a list of components or a component with a specified component identifier.",
	Run: func(cmd *cobra.Command, args []string) {
		apiKeyFromEnv := utils.GetEnv("API_KEY")
		apiKey = utils.DefaultToEnv(apiKeyFromEnv, apiKey)

		if componentID == "" {
			url = "https://api.statuspage.io/v1/pages/" + pageID + "/components"
		} else {
			url = "https://api.statuspage.io/v1/pages/" + pageID + "/components/" + componentID
		}

		request, err := http.NewRequest("GET", url, nil)
		request.Header.Set("Authorization", "OAuth "+apiKey)

		client := &http.Client{Timeout: time.Second * 10}
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

var createComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Create a component.",
	Run: func(cmd *cobra.Command, args []string) {
		apiKeyFromEnv := utils.GetEnv("API_KEY")
		apiKey = utils.DefaultToEnv(apiKeyFromEnv, apiKey)

		allowedComponentStatuses := []string{"operational", "under_maintenance", "degraded_performance", "partial_outage", "major_outage"}

		if (utils.Contains(allowedComponentStatuses, componentStatus)) != true {
			cmd.Help()
			os.Exit(1)
		}

		var jsonStr = []byte(fmt.Sprintf(`{"component": {"name": "%s", "description": "%s", "status": "%s", "showcase": %t}}`, componentName, componentDescription, componentStatus, componentShowcase))
		request, err := http.NewRequest("POST", "https://api.statuspage.io/v1/pages/"+pageID+"/components", bytes.NewBuffer(jsonStr))
		request.Header.Set("Authorization", "OAuth "+apiKey)
		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{Timeout: time.Second * 10}
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

var updateComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Update a component.",
	Run: func(cmd *cobra.Command, args []string) {
		apiKeyFromEnv := utils.GetEnv("API_KEY")
		apiKey = utils.DefaultToEnv(apiKeyFromEnv, apiKey)

		allowedComponentStatuses := []string{"operational", "under_maintenance", "degraded_performance", "partial_outage", "major_outage"}

		if (utils.Contains(allowedComponentStatuses, componentStatus)) != true {
			cmd.Help()
			os.Exit(1)
		}

		var jsonStr = []byte(fmt.Sprintf(`{"component": {"description": "%s", "status": "%s", "showcase": %t}}`, componentDescription, componentStatus, componentShowcase))
		// fmt.Println(string(jsonStr))
		request, err := http.NewRequest("PATCH", "https://api.statuspage.io/v1/pages/"+pageID+"/components/"+componentID, bytes.NewBuffer(jsonStr))
		request.Header.Set("Authorization", "OAuth "+apiKey)
		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{Timeout: time.Second * 10}
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

var deleteComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Delete a component with a specified component identifier.",
	Run: func(cmd *cobra.Command, args []string) {
		apiKeyFromEnv := utils.GetEnv("API_KEY")
		apiKey = utils.DefaultToEnv(apiKeyFromEnv, apiKey)

		request, err := http.NewRequest("DELETE", "https://api.statuspage.io/v1/pages/"+pageID+"/components/"+componentID, nil)
		request.Header.Set("Authorization", "OAuth "+apiKey)

		client := &http.Client{Timeout: time.Second * 10}
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

		fmt.Println(resp.StatusCode, string(body))
	},
}

func init() {
	getComponentCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "API_KEY environment variable. API key to authenticate against the status page API (required)")
	getComponentCmd.Flags().StringVarP(&pageID, "page-id", "p", "", "Page identifier (required)")
	getComponentCmd.Flags().StringVarP(&componentID, "id", "i", "", "Component identifier")
	getComponentCmd.MarkFlagRequired("page-id")
	createComponentCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "API_KEY environment variable. API key to authenticate against the status page API (required)")
	createComponentCmd.Flags().StringVarP(&pageID, "page-id", "p", "", "Page identifier (required)")
	createComponentCmd.Flags().StringVarP(&componentName, "name", "n", "", "Display name for component (required)")
	createComponentCmd.Flags().StringVarP(&componentDescription, "description", "d", "", "More detailed description for component (required)")
	createComponentCmd.Flags().StringVarP(&componentStatus, "status", "s", "", "Status of the component. Valid choices are: operational, under_maintenance, degraded_performance, partial_outage, major_outage (required)")
	createComponentCmd.Flags().BoolVarP(&componentShowcase, "showcase", "c", false, "Should this component be showcased")
	createComponentCmd.MarkFlagRequired("page-id")
	createComponentCmd.MarkFlagRequired("name")
	createComponentCmd.MarkFlagRequired("description")
	createComponentCmd.MarkFlagRequired("status")
	updateComponentCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "API_KEY environment variable. API key to authenticate against the status page API (required)")
	updateComponentCmd.Flags().StringVarP(&pageID, "page-id", "p", "", "Page identifier (required)")
	updateComponentCmd.Flags().StringVarP(&componentID, "id", "i", "", "Component identifier (required)")
	updateComponentCmd.Flags().StringVarP(&componentDescription, "description", "d", "", "More detailed description for component")
	updateComponentCmd.Flags().StringVarP(&componentStatus, "status", "s", "", "Status of the component. Valid choices are: operational, under_maintenance, degraded_performance, partial_outage, major_outage (required)")
	updateComponentCmd.Flags().BoolVarP(&componentShowcase, "showcase", "c", false, "Should this component be showcased")
	updateComponentCmd.MarkFlagRequired("page-id")
	updateComponentCmd.MarkFlagRequired("id")
	updateComponentCmd.MarkFlagRequired("status")
	deleteComponentCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "API_KEY environment variable. API key to authenticate against the status page API (required)")
	deleteComponentCmd.Flags().StringVarP(&componentID, "id", "i", "", "Component identifier (required)")
	deleteComponentCmd.Flags().StringVarP(&pageID, "page-id", "p", "", "Page identifier (required)")
	deleteComponentCmd.MarkFlagRequired("id")
	deleteComponentCmd.MarkFlagRequired("page-id")
	createCmd.AddCommand(createComponentCmd)
	updateCmd.AddCommand(updateComponentCmd)
	deleteCmd.AddCommand(deleteComponentCmd)
	getCmd.AddCommand(getComponentCmd)
}
