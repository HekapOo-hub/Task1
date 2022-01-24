package model

import "fmt"

type Human struct {
	Id   int
	Name string
	Male bool
	Age  int
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
