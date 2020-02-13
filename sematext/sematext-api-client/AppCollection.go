package sematext

import "errors"

// AppCollection TODO Doc Comment
type AppCollection struct {
	Apps []App `json:"apps"`
}

// GetCollection TODO Doc Comment
func (appCollection *AppCollection) GetCollection(config Configuration) error {

	if config.Client.CachedToken == "" {
		return errors.New("Bad or missing token")
	}
	err := config.Client.GetJSON("/users-web/api/v3/apps", appCollection)
	if err != nil {
		return err
	}

	return nil
}
