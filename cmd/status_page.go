package cmd

import (
	"github.com/TwiN/gatus-client/config"
	"github.com/spf13/cobra"
)

var statusPageCmd = &cobra.Command{
	Use:       "status-page",
	Short:     "Manage status pages",
	Aliases:   []string{"sp"},
	ValidArgs: []string{"get"},
}

var statusPageGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a status page",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := config.GetAPIKey()
		_ = apiKey
		// TODO: Send a request to https://gatus.io/api/v1/external/status-pages/{id}
	},
}

func init() {
	rootCmd.AddCommand(statusPageCmd)
	statusPageCmd.AddCommand(statusPageGetCmd)
}
