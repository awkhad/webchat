package controllers

import (
    "github.com/robfig/revel"
    "webchat/app/model"
    "webchat/app/form"
    //"fmt"
)

type Users struct {
	*Application
}

func (c Users) New() revel.Result {
	return c.Render()
}

func (c Users) Create(userform *form.UserForm) revel.Result {

    userform.Validate(c.Validation)

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(Users.New)
    }

    user := model.NewUser(userform)
    err := user.Save()

    if  err != nil {
        c.Flash.Error(err.Error())
        return c.Redirect(Users.New)
    }

    return c.Redirect(Application.Index)
}

