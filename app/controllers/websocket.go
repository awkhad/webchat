package controllers

import (
    "github.com/robfig/revel"
    "code.google.com/p/go.net/websocket"
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
    return nil
}

