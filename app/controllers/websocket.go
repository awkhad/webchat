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

    fmt.Println("the room key is:", roomkey)

    for room := CharServer.Rooms.Front(); room != nil; room = room.Next() {
        r := room.Value.(chatserver.Room)

        if r.RoomKey == roomkey {
            r.JoinUser(c.Session["user_name"])
        }
        //go room run 
    }
    
    return nil
}

