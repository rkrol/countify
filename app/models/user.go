package models

import (
	"fmt"
	"github.com/revel/revel"
)

type User struct {
	UserId			int
	Name				string
	Email				string

	Accounts		[]*Account
}

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)

	// TODO check email
}
