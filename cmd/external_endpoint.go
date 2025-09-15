package cmd

import (
	"context"
	"fmt"

	gatussdk "github.com/TwiN/gatus-sdk"
	"github.com/spf13/cobra"
)

var externalEndpointCmd = &cobra.Command{
	Use:     "external-endpoint",
	Short:   "Interact with external endpoints",
	Aliases: []string{"external-endpoints", "ee"},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Flags().GetString("url")
		if url == "" {
			return fmt.Errorf("missing required flag --url")
		}
		return nil
	},
}

var externalEndpointPushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push monitoring result to an external endpoint in Gatus",
	Long:  `Push a monitoring result to an external endpoint in Gatus. This is used for push-based monitoring where external systems can report their health status to Gatus. The endpoint must be configured as an external endpoint in Gatus with a matching token.`,
	Example: `  gatus-cli external-endpoint push --url https://status.example.com --key "group_endpoint" --token "secret-token" --success
  gatus-cli external-endpoint push --url https://status.example.com --key "group_endpoint" --token "secret-token" --error "Connection timeout" --duration "2s"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Parent().Flags().GetString("url")
		key, _ := cmd.Flags().GetString("key")
		token, _ := cmd.Flags().GetString("token")
		success, _ := cmd.Flags().GetBool("success")
		errorMessage, _ := cmd.Flags().GetString("error")
		duration, _ := cmd.Flags().GetString("duration")

		if key == "" || token == "" {
			return fmt.Errorf("missing required flags --key and --token")
		}

		// If error message is provided, success should be false
		if errorMessage != "" {
			success = false
		}

		gatusClient := gatussdk.NewClient(url)
		err := gatusClient.PushExternalEndpointResult(context.Background(), key, token, success, errorMessage, duration)
		if err != nil {
			return fmt.Errorf("failed to push result: %w", err)
		}

		if success {
			fmt.Printf("successfully pushed healthy status on endpoint with key %s to %s", key, url)
		} else {
			fmt.Printf("successfully pushed unhealthy status on endpoint with key %s to %s", key, url)
		}
		return nil
	},
}

func init() {
	// Add external-endpoint command to root with persistent --url flag
	externalEndpointCmd.PersistentFlags().StringP("url", "u", "", "Gatus server URL (required)")
	rootCmd.AddCommand(externalEndpointCmd)

	// Add push command
	externalEndpointPushCmd.Flags().StringP("key", "k", "", "Endpoint key in format {group}_{name} (required)")
	externalEndpointPushCmd.Flags().StringP("token", "t", "", "Bearer token for external endpoint (required)")
	externalEndpointPushCmd.Flags().BoolP("success", "s", false, "Mark the health check as successful")
	externalEndpointPushCmd.Flags().StringP("error", "e", "", "Error message if check failed")
	externalEndpointPushCmd.Flags().StringP("duration", "d", "", "Duration of the health check (e.g. '10s', '500ms')")
	externalEndpointCmd.AddCommand(externalEndpointPushCmd)
}
