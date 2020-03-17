package sematext

import (
	"bytes"
	"encoding/json"

	"github.com/google/uuid"
)

// IsValidSematextRegion checks sematext api region is valid.
func IsValidSematextRegion(region string) bool {
	switch region {
	case "EU", "US":
		return true
	}
	return false
}

// IsValidUUID checks a string is UUIDv4
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

// PrettyPrintJSON TODO Doc Comment
func PrettyPrintJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}
