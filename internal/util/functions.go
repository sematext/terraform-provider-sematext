package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/sematext/sematext-api-client-go/stcloud"

	"github.com/sematext/terraform-provider-sematext/internal/common"
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
func ExtractApp(appResponse *stcloud.AppResponse) (*stcloud.App, error) {

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
func ExtractFirstApp(appsResponse *stcloud.AppsResponse) (*stcloud.App, error) {

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
func ExtractAppTokens(tokensResponse stcloud.TokensResponse) (*[]stcloud.TokenDto, error) {

	if tokensResponse.Success {
		if tokensResponse.Data != nil {
			if tokensResponse.Data.Tokens != nil {
				return &tokensResponse.Data.Tokens, nil
			}
		}
	}

	return nil, errors.New(tokensResponse.Errors[0].Message)
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ArrayLiteralString(array []string) string {
	s := strings.Join(array, "\",\"")
	return fmt.Sprintf("[\"%s\"]", s)
}

func ExtractAppTokenNames(AppTokens [](common.AppTokenType)) []string {

	var result []string
	for _, token := range AppTokens {
		for _, name := range token.Names {
			result = append(result, name)
		}
	}
	return result
}

func Int64StringMapKeys(m map[int64]string) []int64 {
	i := 0
	keys := make([]int64, len(m))
	for k := range m {
		keys[i] = int64(k)
		i++
	}
	return keys
}

func StringStringMapKeys(m map[string]string) []string {
	i := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[i] = string(k)
		i++
	}
	return keys
}

func IntArraytoInt64array(a []int) []int64 {
	result := make([]int64, len(a))
	for i, v := range a {
		result[i] = int64(v)
	}
	return result

}