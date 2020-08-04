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
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	// "github.com/appvia/statuspage/utils"
	"../utils"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var incidentID string
var incidentName string
var incidentStatus string
var incidentBody string
var incidentComponents map[string]string

var getIncidentCmd = &cobra.Command{
	Use:   "incident",
	Short: "Get a list of incidents or a component with a specified incident identifier.",
	Run: func(cmd *cobra.Command, args []string) {
		if incidentID == "" {
			url = "https://api.statuspage.io/v1/pages/" + pageID + "/incidents"
		} else {
			url = "https://api.statuspage.io/v1/pages/" + pageID + "/incidents/" + incidentID
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

var createIncidentCmd = &cobra.Command{
	Use:   "incident",
	Short: "Create an incident.",
	Run: func(cmd *cobra.Command, args []string) {
		incidentComponentIDs := []string{}
		for c, _ := range incidentComponents {
			incidentComponentIDs = append(incidentComponentIDs, c)
		}

		jsonIncidentComponentIDs, _ := json.Marshal(incidentComponentIDs)
		jsonIncidentComponents, _ := json.Marshal(incidentComponents)

		allowedIncidentStatuses := []string{"investigating", "identified", "monitoring", "resolved", "scheduled", "in_progress", "verifying", "completed"}

		if (utils.Contains(allowedIncidentStatuses, incidentStatus)) != true {
			cmd.Help()
			os.Exit(1)
		}

		var jsonStr = []byte(fmt.Sprintf(`{"incident": {"name": "%s", "status": "%s", "body": "%s", "component_ids": %s, "components": %s}}`, incidentName, incidentStatus, incidentBody, jsonIncidentComponentIDs, jsonIncidentComponents))
		request, err := http.NewRequest("POST", "https://api.statuspage.io/v1/pages/"+pageID+"/incidents", bytes.NewBuffer(jsonStr))
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

var updateIncidentCmd = &cobra.Command{
	Use:   "incident",
	Short: "Update an incident.",
	Run: func(cmd *cobra.Command, args []string) {
		incidentComponentIDs := []string{}
		for c, _ := range incidentComponents {
			incidentComponentIDs = append(incidentComponentIDs, c)
		}

		jsonIncidentComponentIDs, _ := json.Marshal(incidentComponentIDs)
		jsonIncidentComponents, _ := json.Marshal(incidentComponents)

		allowedIncidentStatuses := []string{"investigating", "identified", "monitoring", "resolved", "scheduled", "in_progress", "verifying", "completed"}

		if (utils.Contains(allowedIncidentStatuses, incidentStatus)) != true {
			cmd.Help()
			os.Exit(1)
		}

		var jsonStr = []byte(fmt.Sprintf(`{"incident": {"status": "%s", "body": "%s", "component_ids": %s, "components": %s}}`, incidentStatus, incidentBody, jsonIncidentComponentIDs, jsonIncidentComponents))
		request, err := http.NewRequest("PATCH", "https://api.statuspage.io/v1/pages/"+pageID+"/incidents/"+incidentID, bytes.NewBuffer(jsonStr))
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

var deleteIncidentCmd = &cobra.Command{
	Use:   "incident",
	Short: "Delete an incident with a specified incident identifier.",
	Run: func(cmd *cobra.Command, args []string) {
		request, err := http.NewRequest("DELETE", "https://api.statuspage.io/v1/pages/"+pageID+"/incidents/"+incidentID, nil)
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

func init() {
	getIncidentCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "API key to authenticate against the status page API (required)")
	getIncidentCmd.Flags().StringVarP(&pageID, "page-id", "p", "", "Page identifier (required)")
	getIncidentCmd.Flags().StringVarP(&incidentID, "id", "i", "", "Incident Identifier")
	getIncidentCmd.MarkFlagRequired("api-key")
	getIncidentCmd.MarkFlagRequired("page-id")
	createIncidentCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "API key to authenticate against the status page API (required)")
	createIncidentCmd.Flags().StringVarP(&pageID, "page-id", "p", "", "Page identifier (required)")
	createIncidentCmd.Flags().StringVarP(&incidentName, "name", "n", "", "Incident name (required)")
	createIncidentCmd.Flags().StringVarP(&incidentStatus, "status", "s", "", "The Incident status. Valid choices are: investigating, identified, monitoring, resolved, scheduled, in_progress, verifying, completed.")
	createIncidentCmd.Flags().StringVarP(&incidentBody, "body", "b", "", "The initial message, created as the first incident update")
	createIncidentCmd.Flags().StringToStringVarP(&incidentComponents, "components", "c", map[string]string{}, "Map of status changes to apply to affected components")
	createIncidentCmd.MarkFlagRequired("api-key")
	createIncidentCmd.MarkFlagRequired("page-id")
	createIncidentCmd.MarkFlagRequired("name")
	updateIncidentCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "API key to authenticate against the status page API (required)")
	updateIncidentCmd.Flags().StringVarP(&pageID, "page-id", "p", "", "Page identifier (required)")
	updateIncidentCmd.Flags().StringVarP(&incidentID, "id", "i", "", "Incident identifier (required)")
	updateIncidentCmd.Flags().StringVarP(&incidentStatus, "status", "s", "", "The incident status. Valid choices are: investigating, identified, monitoring, resolved, scheduled, in_progress, verifying, completed.")
	updateIncidentCmd.Flags().StringVarP(&incidentBody, "body", "b", "", "The initial message, created as the first incident update")
	updateIncidentCmd.Flags().StringToStringVarP(&incidentComponents, "components", "c", map[string]string{}, "Map of status changes to apply to affected components")
	updateIncidentCmd.MarkFlagRequired("api-key")
	updateIncidentCmd.MarkFlagRequired("page-id")
	updateIncidentCmd.MarkFlagRequired("id")
	deleteIncidentCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "API key to authenticate against the status page API (required)")
	deleteIncidentCmd.Flags().StringVarP(&incidentID, "id", "i", "", "Incident Identifier (required)")
	deleteIncidentCmd.Flags().StringVarP(&pageID, "page-id", "p", "", "Page identifier (required)")
	deleteIncidentCmd.MarkFlagRequired("api-key")
	deleteIncidentCmd.MarkFlagRequired("id")
	deleteIncidentCmd.MarkFlagRequired("page-id")
	getCmd.AddCommand(getIncidentCmd)
	createCmd.AddCommand(createIncidentCmd)
	updateCmd.AddCommand(updateIncidentCmd)
	deleteCmd.AddCommand(deleteIncidentCmd)
}
