package sybilion

import (
	"encoding/json"
	"errors"
	"fmt"

	api "go.sybilion.dev/sybilion/api"
)

// parseAPIError inspects a GenericOpenAPIError returned by the generated client,
// tries to parse the response body as {"error": "<message>"}, and returns a plain
// error containing just that message. Falls back to the original error untouched.
func parseAPIError(err error) error {
	if err == nil {
		return nil
	}
	var apiErr *api.GenericOpenAPIError
	if errors.As(err, &apiErr) {
		var body struct {
			Error string `json:"error"`
		}
		if json.Unmarshal(apiErr.Body(), &body) == nil && body.Error != "" {
			return fmt.Errorf("%s", body.Error)
		}
	}
	return err
}
