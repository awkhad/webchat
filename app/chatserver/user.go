package chatserver 

import (
    "code.google.com/p/go.net/websocket"
    "webchat/app/model"
    //"container/list" 
    "fmt"
)

type OnlineUser struct {
    Id int
    Connection   *websocket.Conn
    Send chan *Event 
    Room *ActiveRoom
    Info *UserInfo
}

type UserInfo struct {
    Name string
    Email string
}

func NewOnlineUser(user *model.User, ws *websocket.Conn,room *ActiveRoom) *OnlineUser {
    onlineUser := &OnlineUser{
        Id: user.Id, 
        Connection: ws,
        Send: make(chan *Event, 512),
        Room: room,
        Info: &UserInfo{
            Name: user.Name,
            Email: user.Email,
        },
    }
    return onlineUser
}

func (u *OnlineUser) PushToClient(){
    for b := range u.Send {
        err := websocket.JSON.Send(u.Connection, b)
        if err != nil {
            break
        }
    }
}

func (u *OnlineUser) PullFromClient(){
    for{
        var event Event
        err := websocket.JSON.Receive(u.Connection, &event)
        fmt.Println("the message is:", event)

        // user close
        if err != nil {
            fmt.Println("Receive occur some error")
            return
        }
        u.Room.Broadcast <- &event
    }
}

