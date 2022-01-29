package model

import "fmt"

type Human struct {
	ID   string
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
	return fmt.Sprintf("Human's info:\nId:%s\nName:%s\nSex:%s\nAge:%d", h.ID, h.Name, male, h.Age)
}

type User struct {
	ID       string
	Login    string
	Password string
	Role     string
}

func (u User) String() string {
	res := fmt.Sprintf("User's info:\nId:%s\nLogin:%s\nPassword:%s\n", u.ID, u.Login, u.Password)
	if u.Role == "admin" {
		res += "Role:admin"
	}
	return res
}
