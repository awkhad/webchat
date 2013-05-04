package controllers

import (
    "github.com/robfig/revel"
    "code.google.com/p/go.net/websocket"
    "webchat/app/chatserver"
    "fmt"
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
    onlineUser := chatserver.NewOnlineUser(user, ws, activeRoom)
    activeRoom.JoinUser(onlineUser)
    go onlineUser.PushToClient()
    onlineUser.PullFromClient()

    fmt.Println("the room count is:", ChatServer.ActiveRooms.Len())
    return nil
}

