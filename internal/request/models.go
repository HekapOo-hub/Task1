package request

type CreateHumanRequest struct {
	Name string `json:"name" form:"name"`
	Male bool   `json:"male" form:"male"`
	Age  int    `json:"age" form:"age"`
}
type UpdateHumanRequest struct {
	Id      string `json:"id" form:"name"`
	NewName string `json:"name" form:"name"`
	NewMale bool   `json:"male" form:"male"`
	NewAge  int    `json:"age" form:"age"`
}
type UpdateUserRequest struct {
	OldLogin    string `json:"oldLogin"`
	NewLogin    string `json:"newLogin"`
	NewPassword string `json:"password"`
}
