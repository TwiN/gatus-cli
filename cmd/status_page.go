package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

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
	Use:     "get",
	Short:   "Retrieve a status page",
	Example: `  gatus-client status-page get --status-page-id 12345`,
	RunE: func(cmd *cobra.Command, args []string) error {
		statusPageID, _ := cmd.Flags().GetInt("status-page-id")
		if statusPageID == 0 {
			return errors.New("missing flag --status-page-id")
		}
		apiKey := config.GetAPIKey()
		request, err := http.NewRequest("GET", "https://gatus.io/api/v1/external/status-pages/"+strconv.Itoa(statusPageID), nil)
		if err != nil {
			return err
		}
		request.Header.Set("Authorization", "Bearer "+apiKey)
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		if response.StatusCode != http.StatusOK {
			// Return the error message as a string
			return errors.New(string(body))
		}
		// Else, print the body and return no error
		fmt.Printf(string(body))
		return nil
	},
}

func init() {
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
	rootCmd.AddCommand(statusPageCmd)
	statusPageGetCmd.Flags().IntP("status-page-id", "s", 0, "ID of the status page")
	statusPageCmd.AddCommand(statusPageGetCmd)
}
