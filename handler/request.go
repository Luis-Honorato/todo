package handler

import "fmt"

func errMandatoryField(v interface{}, typ string) error {
	return fmt.Errorf("error mandatory field: %v (type-%s)", v, typ)
}

// Create Todo Response

// Expected request when creating todo
type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Finished    *bool  `json:"finished"`
}

// Validation to assert Create Todo Request is valid, don't miss mandatory fiels,
// Or is blanked
func (r *CreateTodoRequest) Validate() error {
	if r.Title == "" && r.Description == "" && r.Finished == nil {
		return fmt.Errorf("error blanked request")
	}

	// Set Title as a mandatory field of requisition
	if r.Title == "" {
		return errMandatoryField(r.Title, "string")

	}

	// Set Description as a mandatory field of requisition
	if r.Description == "" {
		return errMandatoryField(r.Description, "string")

	}

	// Set Finished as a mandatory field of requisition
	if r.Finished == nil {
		return errMandatoryField(r.Finished, "bool")

	}

	// Valid requisition
	return nil
}

// Delete Todo Response

// Expected request when deleting todo
type DeleteTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Finished    *bool  `json:"finished"`
}

// Validate to assert at Delete Todo request is valid and has
// at least one attribute to update
func (r *DeleteTodoRequest) Validate() error {

	// Assert requisition will has at least one attribute
	if r.Title == "" || r.Description == "" || r.Finished == nil {
		return nil

	}

	// Returns a error if don't has any attribute
	return fmt.Errorf("provide at least one change")
}
