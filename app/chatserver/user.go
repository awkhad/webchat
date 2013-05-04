package chatserver 

import (
    "code.google.com/p/go.net/websocket"
    //"container/list" 
    "fmt"
)

type OnlineUser struct {
    Name string
    Connection   *websocket.Conn
    Send chan Event 
    Room *ActiveRoom
}

//type UserInfo struct {
//}

func NewOnlineUser(name string, ws *websocket.Conn,room *ActiveRoom) *OnlineUser {
    user := &OnlineUser{
        Name: name, 
        Connection: ws,
        Send: make(chan Event, 512),
        Room: room,
    }
    return user
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
        //var event string
        //var eventtest Event
        //err := websocket.Message.Receive(u.Connection, &event)
        fmt.Println("the message is:", event)

        // user close
        if err != nil {
            fmt.Println("Receive occur some error")
            return
        }
        u.Room.Broadcast <- event
        //u.Room.Broadcast <- eventtest
    }
}

