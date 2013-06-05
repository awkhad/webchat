package controllers

import (
	"github.com/robfig/revel"
	"webchat/app/model"
)

type Admin struct {
	*Application
}

func (c Admin) Index() revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	if !c.checkAdmin() {
		c.Flash.Error("required admin")
		return c.Redirect(Application.Index)
	}

	roomCount := ChatServer.ActiveRooms.Len()
	onlineUserCount := ChatServer.ActiveUsers.Len()
	latestUsers := model.LatestUsers(5)
	userCount := model.UserCount()

	return c.Render(roomCount, onlineUserCount, latestUsers, userCount)
}

func (c Admin) checkAdmin() bool {
	user := model.FindUserByName(c.Session["user_name"])
	if user.Email == "ldshuang@gmail.com" {
		return true
	} else {
		return false
	}
}
