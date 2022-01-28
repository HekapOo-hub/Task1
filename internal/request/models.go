package request

type CreateRequest struct {
	Name string `json:"name" form:"name"`
	Male bool   `json:"male" form:"male"`
	Age  int    `json:"age" form:"age"`
}
type UpdateRequest struct {
	Id      string `json:"id" form:"name"`
	NewName string `json:"name" form:"name"`
	NewMale bool   `json:"male" form:"male"`
	NewAge  int    `json:"age" form:"age"`
}
