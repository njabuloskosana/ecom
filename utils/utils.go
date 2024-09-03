package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Reusable functions
func ParseJSON(r *http.Request, payload interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(payload)

}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)

}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})

}

// singleton to avid creating multiple instances of the validator
var Validate = validator.New()
