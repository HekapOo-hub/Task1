package model

import "fmt"

type Human struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Male bool   `json:"male" form:"male"`
	Age  int    `json:"age" form:"age"`
}

type Id struct {
	Id int `json:"id" form:"id"`
}

func (h Human) String() string {
	var male string
	if h.Male {
		male = "Male"
	} else {
		male = "Female"
	}
	return fmt.Sprintf("Human's info:\nId:%d\nName:%s\nSex:%s\nAge:%d", h.Id, h.Name, male, h.Age)
}
