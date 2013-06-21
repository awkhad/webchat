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

func (c Admin) Users() revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	if !c.checkAdmin() {
		c.Flash.Error("required admin")
		return c.Redirect(Application.Index)
	}

	users := model.AllUsers()
	return c.Render(users)
}

func (c Admin) Rooms() revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	if !c.checkAdmin() {
		c.Flash.Error("required admin")
		return c.Redirect(Application.Index)
	}

	rooms := ChatServer.AllRunRooms()
	return c.Render(rooms)

}

func (c Admin) ChangeLogStatus(roomKey string) revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	if !c.checkAdmin() {
		c.Flash.Error("required admin")
		return c.Redirect(Application.Index)
	}

    room := ChatServer.GetActiveRoom(roomKey)
    room.SaveLogs = !room.SaveLogs
    return c.Redirect(Admin.Rooms)
}

func (c Admin) checkAdmin() bool {
	user := model.FindUserByName(c.Session["user_name"])
	if user.Email == "ldshuang@gmail.com" {
		return true
	} else {
		return false
	}
}

