package form

import (
	"github.com/robfig/revel"
)

// user signup form field
type UserForm struct {
	Name           string
	Email          string
	Password       string
	RepeatPassword string
}

func (userform *UserForm) Validate(v *revel.Validation) {
	v.Required(userform.Name).Message("please verify you name")
	v.Required(userform.RepeatPassword).Message("please verify you repeat password ")
	v.Required(userform.Password).Message("please verify you password")
	v.Required(userform.Email).Message("please verify you Email")
	v.Email(userform.Email).Message("please verify you Email")
}

type UserLogin struct {
	Name, Password string
	Remember       bool
}

func (loginform *UserLogin) Validate(v *revel.Validation) {
	v.Required(loginform.Name)
	v.Required(loginform.Password)
}
