// Package request is a package of structs for binding json data from request
package request

// CreateHumanRequest is a struct to which I bind json data from create human request
type CreateHumanRequest struct {
	Name string `json:"name" validate:"required"`
	Male bool   `json:"male" validate:"required"`
	Age  int    `json:"age" validate:"required"`
}

// UpdateHumanRequest is a struct to which I bind json data from update human request
type UpdateHumanRequest struct {
	ID      string `json:"id" validate:"required"`
	NewName string `json:"name" validate:"required"`
	NewMale bool   `json:"male" validate:"required"`
	NewAge  int    `json:"age" validate:"required"`
}

// UpdateUserRequest is a struct to which I bind json data from update user request
type UpdateUserRequest struct {
	OldLogin    string `json:"oldLogin" validate:"required"`
	NewLogin    string `json:"newLogin" validate:"required"`
	NewPassword string `json:"password" validate:"required"`
}
