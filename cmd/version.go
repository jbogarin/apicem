package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gives the version",
	Long:  `Gives the version of the client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 0.0.1")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
