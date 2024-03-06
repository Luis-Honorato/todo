package handler

import "fmt"

func errMandatoryField(v interface{}, typ string) error {
	return fmt.Errorf("error mandatory field: %v (type-%s)", v, typ)
}

// Create Todo Response

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Finished    *bool  `json:"finished"`
}

func (r *CreateTodoRequest) Validate() error {
	if r.Title == "" {
		return errMandatoryField(r.Title, "string")

	}
	if r.Description == "" {
		return errMandatoryField(r.Description, "string")

	}
	if r.Finished == nil {
		return errMandatoryField(r.Finished, "bool")

	}

	return nil
}

// Delete Todo Response

type DeleteTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Finished    *bool  `json:"finished"`
}

func (r *DeleteTodoRequest) Validate() error {
	if r.Title == "" || r.Description == "" || r.Finished == nil {
		return nil

	}

	return fmt.Errorf("provide at least one change")
}
