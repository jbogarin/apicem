package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	apicem "github.com/jbogarin/go-apic-em/apic-em"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var username string
var password string

// ticketCmd represents the ticket command
var ticketCmd = &cobra.Command{
	Use:   "ticket",
	Short: "Gets the ticket",
	Long:  `Authenticates with the API and gets the ticket`,
	Run: func(cmd *cobra.Command, args []string) {
		user := &apicem.User{
			Username: username,
			Password: password,
		}

		newTicket, _, err := Client.Ticket.AddTicket(user)
		if err != nil {
			log.Fatal(err)
		}

		ticketJSON, err := json.MarshalIndent(newTicket, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(ticketJSON))
		os.Setenv("CISCO_APIC-EM_TICKET", newTicket.Response.ServiceTicket)
		viper.Set("ticket", newTicket.Response.ServiceTicket)

	},
}

func init() {
	RootCmd.AddCommand(ticketCmd)

	ticketCmd.Flags().StringVarP(&username, "username", "u", "", "username")
	ticketCmd.Flags().StringVarP(&password, "password", "p", "", "password")

}
