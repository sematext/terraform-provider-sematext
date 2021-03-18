package sematext

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/sematext/sematext-api-client-go/stcloud"
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

// PrettyPrintJSON is a utility function to format a JSON byte array.
func PrettyPrintJSON(b []byte) ([]byte, error) {

	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")

	return out.Bytes(), err
}

// ExtractApp - cope with missing field or empty array in appsResponse.Data
func extractApp(appResponse *stcloud.AppResponse) (*stcloud.App, error) {

	if appResponse.Success {
		if appResponse.Data != nil {
			if appResponse.Data.App != nil {
				return appResponse.Data.App, nil
			}
		}
	}

	return nil, errors.New(appResponse.Errors[0].Message)
}

// ExtractApp - cope with missing field or empty array in appsResponse.Data
func extractFirstApp(appsResponse *stcloud.AppsResponse) (*stcloud.App, error) {

	if appsResponse.Success {
		if appsResponse.Data != nil {
			if appsResponse.Data.Apps != nil {
				return &appsResponse.Data.Apps[0], nil
			}
		}
	}

	return nil, errors.New(appsResponse.Errors[0].Message)
}

// ExtractAppTokens - cope with missing field or empty array in Response to token operation
func extractAppTokens(tokensResponse stcloud.TokensResponse) (*[]stcloud.TokenDto, error) {

	if tokensResponse.Success {
		if tokensResponse.Data != nil {
			if tokensResponse.Data.Tokens != nil {
				return &tokensResponse.Data.Tokens, nil
			}
		}
	}

	return nil, errors.New(tokensResponse.Errors[0].Message)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
