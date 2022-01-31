package request

type CreateHumanRequest struct {
	Token string `json:"token"`
	Name  string `json:"name" form:"name"`
	Male  bool   `json:"male" form:"male"`
	Age   int    `json:"age" form:"age"`
}
type UpdateHumanRequest struct {
	Token   string `json:"token"`
	Id      string `json:"id" form:"id"`
	NewName string `json:"name" form:"name"`
	NewMale bool   `json:"male" form:"male"`
	NewAge  int    `json:"age" form:"age"`
}
type CreateUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UpdateUserRequest struct {
	Token       string `json:"token"`
	OldLogin    string `json:"oldLogin"`
	NewLogin    string `json:"newLogin"`
	NewPassword string `json:"password"`
}
