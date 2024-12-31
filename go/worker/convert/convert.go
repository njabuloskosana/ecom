package convert

import (
	"log"
	"strconv"
)

// MustConvertToInt converts a string to an integer and logs a fatal error if the conversion fails.
func MustConvertToInt(s string) int {
	port, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}
	return port
}
