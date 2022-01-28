package model

import "fmt"

type Human struct {
	Id   string
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
	return fmt.Sprintf("Human's info:\nId:%s\nName:%s\nSex:%s\nAge:%d", h.Id, h.Name, male, h.Age)
}

type User struct {
	Id       string
	Login    string
	Password string
	Role     string
}

func (u User) String() string {
	res := fmt.Sprintf("User's info:\nId:%s\nLogin:%s\nPassword:%s\n", u.Id, u.Login, u.Password)
	if u.Role == "admin" {
		res += "Role:admin"
	}
	return res
}
