package sematext

import "log"

// AppCollection TODO Doc Comment
type AppCollection struct {
	Apps []App `json:"apps"` // TODO Check datatype and json tag are correct.
}

// Get TODO Doc Comment
func (appCollection AppCollection) Get(token string) (AppCollection, error) {

	err := c.GetJSON("/users-web/api/v3/apps", appCollection) // TODO - unrolling an array of App structs
	if err != nil {
		log.Fatal(err)
	}
	// TODO Forall applicationCollection into a map by id
	return appCollection, nil
}
