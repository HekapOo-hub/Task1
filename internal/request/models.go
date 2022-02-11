// Package request is a package of structs for binding json data from request
package request

// CreateHumanRequest is a struct to which I bind json data from create human request
type CreateHumanRequest struct {
	Name string `json:"name" validate:"required,max=40"`
	Male bool   `json:"male"`
	Age  int    `json:"age" validate:"required"`
}

// UpdateHumanRequest is a struct to which I bind json data from update human request
type UpdateHumanRequest struct {
	OldName string `json:"oldName" validate:"required"`
	NewName string `json:"name" validate:"required,max=40"`
	NewMale bool   `json:"male"`
	NewAge  int    `json:"age" validate:"required"`
}

// UpdateUserRequest is a struct to which I bind json data from update user request
type UpdateUserRequest struct {
	OldLogin    string `json:"oldLogin" validate:"required"`
	NewLogin    string `json:"newLogin" validate:"required"`
	NewPassword string `json:"password" validate:"required"`
}

// SignInRequest is used for validating sign in data request
type SignInRequest struct {
	Login    string `validate:"required"`
	Password string `validate:"required"`
}

// CreateUserRequest is used for validating create user data in request
type CreateUserRequest struct {
	Login      string `json:"login" validate:"required"`
	Password   string `json:"password" validate:"required"`
	RePassword string `json:"re_password" validate:"required,eqfield=Password"`
}

// GetUserRequest is used for validating get user data in request
type GetUserRequest struct {
	Login string `validate:"required"`
}

// DeleteUserRequest is used for validating delete user data in request
type DeleteUserRequest struct {
	Login string `json:"login" validate:"required"`
}

// GetHumanRequest is used for validating get human data in request
type GetHumanRequest struct {
	Name string `validate:"required,max=40"`
}

// DeleteHumanRequest is used for validating delete human data in request
type DeleteHumanRequest struct {
	Name string `validate:"required"`
}
