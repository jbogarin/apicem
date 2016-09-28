package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"

	apicem "github.com/jbogarin/go-apic-em/apic-em"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var shortResponse bool
var ticketFlag string

// Client is the client
var Client *apicem.Client

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "apicem",
	Short: "Cisco APIC-EM CLI",
	Long:  `Cisco APIC-EM CLI client.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		Client = apicem.NewClient(client)
		Client.BaseURL, _ = url.Parse(viper.GetString("url"))
		if cmd.Use != "ticket" {
			viper.BindEnv("CISCO_APICEM_TICKET")
			var ticket string
			if ticketFlag != "" {
				ticket = ticketFlag
			} else if viper.IsSet("CISCO_APICEM_TICKET") {
				ticket = viper.GetString("CISCO_APICEM_TICKET")
			} else if viper.IsSet("ticket") {
				ticket = viper.GetString("ticket")
			} else {
				fmt.Println(cmd.Help())
				os.Exit(-1)
			}
			Client.Authorization = ticket
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.apicem.yaml)")
	RootCmd.PersistentFlags().StringVarP(&ticketFlag, "ticket", "T", "", "APIC-EM Authorization ticket")
	RootCmd.PersistentFlags().BoolVarP(&shortResponse, "short", "S", false, "Short response")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".apicem") // name of config file (without extension)
	viper.AddConfigPath("$HOME")   // adding home directory as first search path
	viper.AutomaticEnv()           // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
