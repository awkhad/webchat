package controllers

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/robfig/revel"
	"webchat/app/chatserver"
)

type Websocket struct {
	*revel.Controller
}

func (c Websocket) Chat(roomkey string, ws *websocket.Conn) revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	user := CurrentUser(c.Controller)
	activeRoom := ChatServer.GetActiveRoom(roomkey)
	// crate a user and add usr to room
	onlineUser := chatserver.NewOnlineUser(user, ws, activeRoom)
	activeRoom.JoinUser(onlineUser)
	// 
	go onlineUser.PushToClient()
	onlineUser.PullFromClient()

	fmt.Println("the room count is:", ChatServer.ActiveRooms.Len())
	//defer close(onlineUser.Send)
	defer onlineUser.Close()

	return nil
}
