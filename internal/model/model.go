package model

import "fmt"

type Human struct {
	Id   string `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Male bool   `json:"male" form:"male"`
	Age  int    `json:"age" form:"age"`
}

type Name struct {
	Name string `json:"name" form:"name"`
}
type Id struct {
	Id string `json:"id" form:"id"`
}

func (h Human) String() string {
	var male string
	if h.Male {
		male = "Male"
	} else {
		male = "Female"
	}
	return fmt.Sprintf("Human's info:\nId:%s\nName:%s\nSex:%s\nAge:%d", h.Id, h.Name, male, h.Age)
}
