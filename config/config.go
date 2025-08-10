package config

import "os"

func GetAPIKey() string {
	apiKey := os.Getenv("GATUS_CLI_API_KEY")
	if len(apiKey) == 0 {
		panic("GATUS_CLI_API_KEY environment variable not set")
	}
	return apiKey
}
