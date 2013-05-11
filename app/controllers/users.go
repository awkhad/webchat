package controllers

import (
	//"fmt"
	"github.com/robfig/revel"
	"webchat/app/form"
	"webchat/app/model"
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

func (c Users) MyRooms() revel.Result {
	user := CurrentUser(c.Controller)

	rooms := user.Rooms()

	return c.Render(rooms)
}

func (c Users) EditSettings() revel.Result {
	user := CurrentUser(c.Controller)
	return c.Render(user)
}

func (c Users) SaveSettings(setting *form.Settings) revel.Result {
	user := CurrentUser(c.Controller)

	if err := user.SaveSettings(setting); err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(Users.EditSettings)
	}

	c.Flash.Success("save success")
	return c.Redirect(Users.EditSettings)
}

func (c Users) Profile() revel.Result {
	return nil
}

func (c Users) Avatar() revel.Result {
	return nil
}
