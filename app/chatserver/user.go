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

func (u *OnlineUser) Close(){
    // clear resource when user conn close 
    fmt.Println("the user conn is closed...")
    // remove user form rooms 
    for e := u.Room.Users.Front(); e != nil; e = e.Next() {
        user := e.Value.(*OnlineUser)
        if user.Id == u.Id {
            u.Room.Users.Remove(e)
            break
        }
    }

    // close channel
    close(u.Send)
    // send levae message to other client
    event := &Event{
        Type: "leave",
        Text: u.Info.Name + " has leave room",
        User: u.Info,
    }

    u.Room.Broadcast <- event
}

