package sematext

import {
	"errors"
}

// GenericAPIResponse TODO Doc Comment
type GenericAPIResponse struct {
	Data    interface{} `json:"data"`
	Apps    []App       `json:"apps,omitempty"`
	App     App         `json:"apps,omitempty"`
	Errors  []Error     `json:"errors"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

// TODO JIRA TASK - GET App by ID
// TODO JIRA TASK - DELETE App by ID
// TODO JIRA TASK - RESET App STATUS to archived or similar.

func (genericAPIResponse *GenericAPIResponse) extractAppById(id string) (*App, error) {

	if genericAPIResponse.Apps == nil {
		return nil, fmt.Errorf("Missing Apps field")
	}

	if len(genericAPIResponse.Apps) == 0 {
		return nil, nil
	}

	for i := range genericAPIResponse.Apps {
		if genericAPIResponse.Apps[i].ID == id {
			app := genericAPIResponse.Apps[i]
			return &app, nil
		}
	}

	return nil, fmt.Errorf("App %s not found", id)

}

func (genericAPIResponse *GenericAPIResponse) extractApp(id string) (*App, error) {

	if genericAPIResponse.App == nil {
		return nil, fmt.Errorf("Missing App field")
	}

	app := genericAPIResponse.App

	return &app, nil

}
