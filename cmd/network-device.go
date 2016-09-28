package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var networkDeviceScope string

// networkDeviceCmd represents the network-device command
var networkDeviceCmd = &cobra.Command{
	Use:   "network-device",
	Short: "network-device",
	Long:  `network-device`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

// networkDeviceListCmd represents the GET network-device command
var networkDeviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  `list`,
	Run: func(cmd *cobra.Command, args []string) {
		networkDevices, _, err := Client.NetworkDevice.GetAllNetworkDevice(networkDeviceScope)
		if err != nil {
			log.Fatal(err)
		}

		networkDevicesJSON, err := json.MarshalIndent(networkDevices, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(networkDevicesJSON))
	},
}

// networkDeviceCountCmd represents the GET network-device/count command
var networkDeviceCountCmd = &cobra.Command{
	Use:   "count",
	Short: "count",
	Long:  `count`,
	Run: func(cmd *cobra.Command, args []string) {
		networkDevicesCount, _, err := Client.NetworkDevice.GetNetworkDeviceCount(networkDeviceScope)
		if err != nil {
			log.Fatal(err)
		}

		networkDevicesCountJSON, err := json.MarshalIndent(networkDevicesCount, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(networkDevicesCountJSON))

	},
}

func init() {
	RootCmd.AddCommand(networkDeviceCmd)
	networkDeviceCmd.AddCommand(networkDeviceListCmd)
	networkDeviceCmd.AddCommand(networkDeviceCountCmd)
	networkDeviceCmd.PersistentFlags().StringVarP(&networkDeviceScope, "scope", "s", "", "scope")
}
