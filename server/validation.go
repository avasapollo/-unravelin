package server

import "fmt"

type validation struct {
}

func NewValidation() Validation {
	return &validation{}
}

func (validation) ValidateFormRequest(request map[string]interface{}) error {
	if request == nil || len(request) == 0 {
		return fmt.Errorf("the body of the request can not be empty")
	}

	// TODO: I would like to add some validation like sessionID
	return nil
}
