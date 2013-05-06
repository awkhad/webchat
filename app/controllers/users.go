package controllers

import (
	"github.com/robfig/revel"
	"webchat/app/form"
	"webchat/app/model"
	"fmt"
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

	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(Users.New)
	}

	return c.Redirect(Application.Index)
}

func (c Users) EditSettings() revel.Result {
    return c.Render()
}

func (c Users) SaveSettings() revel.Result {
    return nil
}

func (c Users) MyRooms() revel.Result {
    user := CurrentUser(c.Controller)

    fmt.Println(user)

    rooms := user.Rooms()

    return c.Render(rooms)
}
