package controllers

import (
	"github.com/robfig/revel"
)

type Admin struct {
	*revel.Controller
}

func (c Admin) Index() revel.Result {
    roomCount := ChatServer.ActiveRooms.Len()
    userCount := ChatServer.ActiveUsers.Len()
	return c.Render(roomCount, userCount)
}
