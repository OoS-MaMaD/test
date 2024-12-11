package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "watchtower",
	Short: `Usage:
	watchtower-client [flags]`,
	Long: `Usage:
  watchtower-client [flags]

Flags:
  -login      				Authenticate and store tokens
  -programs   				Retrieve programs data
  -http [program_number] 		Retrieve subdomains of program_number
  -h          				Display help`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
