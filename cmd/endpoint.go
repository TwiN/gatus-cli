package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	gatussdk "github.com/TwiN/gatus-sdk"
	"github.com/spf13/cobra"
)

// convertDuration converts Go duration format to Gatus API format
func convertDuration(durationStr string) (string, error) {
	// Valid Gatus durations: 1h, 24h, 7d, 30d
	validDurations := map[string]bool{"1h": true, "24h": true, "7d": true, "30d": true}

	if validDurations[durationStr] {
		return durationStr, nil
	}

	// Map common durations to valid ones
	mappings := map[string]string{
		"1hour": "1h", "hour": "1h",
		"1day": "24h", "day": "24h", "24hour": "24h",
		"week": "7d", "7day": "7d", "7days": "7d",
		"month": "30d", "30day": "30d", "30days": "30d",
	}

	if mapped, exists := mappings[durationStr]; exists {
		return mapped, nil
	}

	return "", fmt.Errorf("invalid duration '%s'. Valid options: 1h, 24h, 7d, 30d", durationStr)
}

var endpointCmd = &cobra.Command{
	Use:     "endpoint",
	Short:   "Interact with endpoints",
	Aliases: []string{"endpoints", "ep"},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Flags().GetString("url")
		if url == "" {
			return fmt.Errorf("missing required flag --url")
		}
		return nil
	},
}

var endpointStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get endpoint status information",
}

var endpointStatusAllCmd = &cobra.Command{
	Use:     "all",
	Short:   "Get status of all endpoints",
	Example: `  gatus-client endpoint status all --url https://status.example.com`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Parent().Parent().Flags().GetString("url")
		gatusClient := gatussdk.NewClient(url)

		statuses, err := gatusClient.GetAllEndpointStatuses(context.Background())
		if err != nil {
			return err
		}

		output, err := json.MarshalIndent(statuses, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
		return nil
	},
}

var endpointStatusGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get status of a specific endpoint by key or by group and name",
	Example: `  gatus-client endpoint status get --url https://status.example.com --key "group_endpoint"
  gatus-client endpoint status get --url https://status.example.com --group "web" --name "frontend"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Parent().Parent().Flags().GetString("url")
		key, _ := cmd.Flags().GetString("key")
		group, _ := cmd.Flags().GetString("group")
		name, _ := cmd.Flags().GetString("name")

		gatusClient := gatussdk.NewClient(url)

		var status *gatussdk.EndpointStatus
		var err error

		if key != "" {
			// Use key if provided
			if group != "" || name != "" {
				return fmt.Errorf("cannot use --key with --group or --name flags")
			}
			status, err = gatusClient.GetEndpointStatusByKey(context.Background(), key)
		} else if group != "" && name != "" {
			// Use group and name if both provided
			status, err = gatusClient.GetEndpointStatus(context.Background(), group, name)
		} else {
			return fmt.Errorf("must provide either --key or both --group and --name")
		}

		if err != nil {
			return err
		}

		output, err := json.MarshalIndent(status, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
		return nil
	},
}

var endpointUptimeCmd = &cobra.Command{
	Use:     "uptime",
	Short:   "Get uptime information for an endpoint",
	Example: `  gatus-client endpoint uptime --url https://status.example.com --key "group_endpoint" --duration "7d"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Parent().Flags().GetString("url")
		key, _ := cmd.Flags().GetString("key")
		durationStr, _ := cmd.Flags().GetString("duration")
		if key == "" || durationStr == "" {
			return fmt.Errorf("missing required flags --key and --duration")
		}

		gatusDuration, err := convertDuration(durationStr)
		if err != nil {
			return err
		}

		gatusClient := gatussdk.NewClient(url)
		uptime, err := gatusClient.GetEndpointUptime(context.Background(), key, gatusDuration)
		if err != nil {
			return err
		}

		fmt.Printf("Uptime: %.2f%%\n", uptime*100)
		return nil
	},
}

var endpointResponseTimesCmd = &cobra.Command{
	Use:     "response-times",
	Short:   "Get response times for an endpoint",
	Example: `  gatus-client endpoint response-times --url https://status.example.com --key "group_endpoint" --duration "24h"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Parent().Flags().GetString("url")
		key, _ := cmd.Flags().GetString("key")
		durationStr, _ := cmd.Flags().GetString("duration")
		if key == "" || durationStr == "" {
			return fmt.Errorf("missing required flags --key and --duration")
		}

		gatusDuration, err := convertDuration(durationStr)
		if err != nil {
			return err
		}

		gatusClient := gatussdk.NewClient(url)
		responseTimes, err := gatusClient.GetEndpointResponseTimes(context.Background(), key, gatusDuration)
		if err != nil {
			return err
		}

		output, err := json.MarshalIndent(responseTimes, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
		return nil
	},
}

var badgeCmd = &cobra.Command{
	Use:   "badge",
	Short: "Generate badge URLs for endpoints",
}

var badgeHealthCmd = &cobra.Command{
	Use:     "health",
	Short:   "Generate health badge URL",
	Example: `  gatus-client endpoint badge health --url https://status.example.com --key "group_endpoint"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Parent().Parent().Flags().GetString("url")
		key, _ := cmd.Flags().GetString("key")
		if key == "" {
			return fmt.Errorf("missing required flag --key")
		}

		gatusClient := gatussdk.NewClient(url)
		badgeURL := gatusClient.GetEndpointHealthBadgeURL(key)
		fmt.Println(badgeURL)
		return nil
	},
}

var badgeUptimeCmd = &cobra.Command{
	Use:     "uptime",
	Short:   "Generate uptime badge URL",
	Example: `  gatus-client endpoint badge uptime --url https://status.example.com --key "group_endpoint" --duration "7d"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Parent().Parent().Flags().GetString("url")
		key, _ := cmd.Flags().GetString("key")
		durationStr, _ := cmd.Flags().GetString("duration")
		if key == "" || durationStr == "" {
			return fmt.Errorf("missing required flags --key and --duration")
		}

		gatusDuration, err := convertDuration(durationStr)
		if err != nil {
			return err
		}

		gatusClient := gatussdk.NewClient(url)
		badgeURL := gatusClient.GetEndpointUptimeBadgeURL(key, gatusDuration)
		fmt.Println(badgeURL)
		return nil
	},
}

var badgeResponseTimeCmd = &cobra.Command{
	Use:     "response-time",
	Short:   "Generate response time badge URL",
	Example: `  gatus-client endpoint badge response-time --url https://status.example.com --key "group_endpoint" --duration "24h"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Parent().Parent().Flags().GetString("url")
		key, _ := cmd.Flags().GetString("key")
		durationStr, _ := cmd.Flags().GetString("duration")
		if key == "" || durationStr == "" {
			return fmt.Errorf("missing required flags --key and --duration")
		}

		gatusDuration, err := convertDuration(durationStr)
		if err != nil {
			return err
		}

		gatusClient := gatussdk.NewClient(url)
		badgeURL := gatusClient.GetEndpointResponseTimeBadgeURL(key, gatusDuration)
		fmt.Println(badgeURL)
		return nil
	},
}

func init() {
	// Add endpoint command to root with persistent --url flag
	endpointCmd.PersistentFlags().StringP("url", "u", "", "Gatus server URL (required)")
	rootCmd.AddCommand(endpointCmd)

	// Add status subcommand
	endpointCmd.AddCommand(endpointStatusCmd)
	endpointStatusCmd.AddCommand(endpointStatusAllCmd)

	endpointStatusGetCmd.Flags().StringP("key", "k", "", "Endpoint key (optional if group/name provided)")
	endpointStatusGetCmd.Flags().StringP("group", "g", "", "Endpoint group (optional if key provided)")
	endpointStatusGetCmd.Flags().StringP("name", "n", "", "Endpoint name (optional if key provided)")
	endpointStatusCmd.AddCommand(endpointStatusGetCmd)

	// Add uptime command
	endpointUptimeCmd.Flags().StringP("key", "k", "", "Endpoint key (required)")
	endpointUptimeCmd.Flags().StringP("duration", "d", "", "Duration: 1h, 24h, 7d, or 30d (required)")
	endpointCmd.AddCommand(endpointUptimeCmd)

	// Add response-times command
	endpointResponseTimesCmd.Flags().StringP("key", "k", "", "Endpoint key (required)")
	endpointResponseTimesCmd.Flags().StringP("duration", "d", "", "Duration: 1h, 24h, 7d, or 30d (required)")
	endpointCmd.AddCommand(endpointResponseTimesCmd)

	// Add badge commands
	endpointCmd.AddCommand(badgeCmd)

	badgeHealthCmd.Flags().StringP("key", "k", "", "Endpoint key (required)")
	badgeCmd.AddCommand(badgeHealthCmd)

	badgeUptimeCmd.Flags().StringP("key", "k", "", "Endpoint key (required)")
	badgeUptimeCmd.Flags().StringP("duration", "d", "", "Duration: 1h, 24h, 7d, or 30d (required)")
	badgeCmd.AddCommand(badgeUptimeCmd)

	badgeResponseTimeCmd.Flags().StringP("key", "k", "", "Endpoint key (required)")
	badgeResponseTimeCmd.Flags().StringP("duration", "d", "", "Duration: 1h, 24h, 7d, or 30d (required)")
	badgeCmd.AddCommand(badgeResponseTimeCmd)
}
