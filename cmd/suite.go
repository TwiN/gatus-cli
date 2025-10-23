package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	gatussdk "github.com/TwiN/gatus-sdk"
	"github.com/spf13/cobra"
)

var suiteCmd = &cobra.Command{
	Use:     "suite",
	Short:   "Interact with suites",
	Aliases: []string{"suites"},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Flags().GetString("url")
		if url == "" {
			return fmt.Errorf("missing required flag --url")
		}
		return nil
	},
}

var suiteStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get suite status information",
}

var suiteStatusAllCmd = &cobra.Command{
	Use:     "all",
	Short:   "Get status of all suites",
	Example: `  gatus-cli suite status all --url https://status.example.com`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Parent().Parent().Flags().GetString("url")
		gatusClient := gatussdk.NewClient(url)
		statuses, err := gatusClient.GetAllSuiteStatuses(context.Background())
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

var suiteStatusGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get status of a specific suite by key or by group and name",
	Example: `  gatus-cli suite status get --url https://status.example.com --key "_check-authentication"
  gatus-cli suite status get --url https://status.example.com --group "monitoring" --name "health-checks"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Parent().Parent().Flags().GetString("url")
		key, _ := cmd.Flags().GetString("key")
		group, _ := cmd.Flags().GetString("group")
		name, _ := cmd.Flags().GetString("name")
		gatusClient := gatussdk.NewClient(url)
		var status *gatussdk.SuiteStatus
		var err error
		if key != "" {
			// Use key if provided
			if group != "" || name != "" {
				return fmt.Errorf("cannot use --key with --group or --name flags")
			}
			status, err = gatusClient.GetSuiteStatusByKey(context.Background(), key)
		} else if name != "" {
			// Use group and name (group is optional for suites)
			status, err = gatusClient.GetSuiteStatus(context.Background(), group, name)
		} else {
			return fmt.Errorf("must provide either --key or --name (with optional --group)")
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

func init() {
	// Add suite command to root with persistent --url flag
	suiteCmd.PersistentFlags().StringP("url", "u", "", "Gatus server URL (required)")
	rootCmd.AddCommand(suiteCmd)
	// Add status subcommand
	suiteCmd.AddCommand(suiteStatusCmd)
	suiteStatusCmd.AddCommand(suiteStatusAllCmd)
	suiteStatusGetCmd.Flags().StringP("key", "k", "", "Suite key (optional if name provided)")
	suiteStatusGetCmd.Flags().StringP("group", "g", "", "Suite group (optional)")
	suiteStatusGetCmd.Flags().StringP("name", "n", "", "Suite name (optional if key provided)")
	suiteStatusCmd.AddCommand(suiteStatusGetCmd)
}
