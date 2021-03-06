// Package model contains different structs which describe info stored in repository
package model

import (
	"fmt"
)

// Human contains data which is stored in postgresRepository
type Human struct {
	ID   string
	Name string
	Male bool
	Age  int
}

// String returns info about human
func (h Human) String() string {
	var male string
	if h.Male {
		male = "Male"
	} else {
		male = "Female"
	}
	return fmt.Sprintf("Human's info:\nId:%s\nName:%s\nSex:%s\nAge:%d", h.ID, h.Name, male, h.Age)
}

// User contains data which is stored in userRepository
type User struct {
	ID       string
	Login    string
	Password string
	Role     string
}

// String returns info about user
func (u User) String() string {
	res := fmt.Sprintf("User's info:\nId:%s\nLogin:%s\nPassword:%s\n", u.ID, u.Login, u.Password)
	if u.Role == "admin" {
		res += "Role:admin"
	}
	return res
}

// Token contains data about refresh token which is stored in tokenRepository
type Token struct {
	Value     string
	ExpiresAt int64
	Login     string
}
