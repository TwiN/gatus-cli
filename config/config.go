package config

import "os"

func GetAPIKey() string {
	apiKey := os.Getenv("GATUS_CLIENT_API_KEY")
	if len(apiKey) == 0 {
		panic("GATUS_CLIENT_API_KEY environment variable not set")
	}
	return apiKey
}
