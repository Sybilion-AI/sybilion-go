package sybilion

import (
	"os"
	"strings"
)

// EnvSybilionAPIBaseURL is the environment variable that overrides the
// compiled-in default API origin when Options.BaseURL is empty.
const EnvSybilionAPIBaseURL = "SYBILION_API_BASE_URL"

// EnvSybilionAPIToken is the environment variable read when Options.Token is empty.
const EnvSybilionAPIToken = "SYBILION_API_TOKEN"

func resolveAPIBaseURL(explicit string) string {
	s := strings.TrimSuffix(strings.TrimSpace(explicit), "/")
	if s != "" {
		return s
	}
	if v := strings.TrimSpace(os.Getenv(EnvSybilionAPIBaseURL)); v != "" {
		return strings.TrimSuffix(v, "/")
	}
	return DefaultPublicAPIBaseURL
}
