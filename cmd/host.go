// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"encoding/json"
	"fmt"
	"log"

	apicem "github.com/jbogarin/go-apic-em/apic-em"
	"github.com/spf13/cobra"
)

var hostsLimit string
var hostsScope string

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "host",
	Long:  `host.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

// hostCmd represents the host command
var hostListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the hosts",
	Long:  `List the hosts using the APIC-EM API.`,
	Run: func(cmd *cobra.Command, args []string) {
		getHostsQueryParams := &apicem.GetHostsQueryParams{
			Limit: hostsLimit,
		}
		hosts, _, err := Client.Host.GetHosts(hostsScope, getHostsQueryParams)
		if err != nil {
			log.Fatal(err)
		}
		hostsJSON, err := json.MarshalIndent(hosts, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(hostsJSON))
	},
}

func init() {
	RootCmd.AddCommand(hostCmd)
	hostCmd.AddCommand(hostListCmd)
	hostListCmd.Flags().StringVarP(&hostsLimit, "max", "m", "", "max items")
	hostListCmd.Flags().StringVarP(&hostsScope, "scope", "s", "", "query scope")
}
