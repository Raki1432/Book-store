package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads the body of an HTTP request and unmarshals it into the specified structure.
func ParseBody(r *http.Request, x interface{}) error {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err // Return the error if body reading fails
	}

	// Unmarshal the JSON into the provided interface
	if err := json.Unmarshal(body, x); err != nil {
		return err // Return the error if unmarshalling fails
	}

	return nil
}
