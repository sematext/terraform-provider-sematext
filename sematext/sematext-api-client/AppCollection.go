package sematext

import "log"

type AppCollection struct {
	Apps []App `json:"apps"` // TODO Check datatype and json tag are correct.
}

func (appCollection AppCollection) Get(token string) (AppCollection, error) {

	err := c.GetJSON("/users-web/api/v3/apps", applicationCollection) // TODO - unrolling an array of App structs
	if err != "" {
		log.Fatal(err)
	}
	// TODO Forall applicationCollection into a map by id
}
