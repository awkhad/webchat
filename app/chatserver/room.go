package chatserver 

import (
    //"code.google.com/p/go.net/websocket"
    "container/list" 
    "fmt"
)

type ActiveRoom struct {
    RoomKey string
    Users   *list.List
    Broadcast chan *Event
}

func NewActiveRoom(rk string) *ActiveRoom{
    activeRoom := &ActiveRoom{
        RoomKey: rk,
        Users: list.New(),
        Broadcast: make(chan *Event),
    }
    return activeRoom
}

func (r ActiveRoom) JoinUser(user *OnlineUser) {
    r.Users.PushBack(user)
    // send join message
    event := &Event{
        Type: "join",
        Text: user.Info.Name + " has join room",
        User: user.Info,
    }

    fmt.Println("the room len is:", r.Users.Len()) 

    r.Broadcast <- event
}

func (r ActiveRoom) UserList() ([]*UserInfo){
    var userList []*UserInfo

    for u := r.Users.Front(); u != nil; u = u.Next() {
        user := u.Value.(*OnlineUser)
        userList = append(userList, user.Info)
    }
    return userList
}

func (r ActiveRoom) Run() {
    for{
        select { 
        case bc := <- r.Broadcast:
            for u := r.Users.Front(); u != nil; u = u.Next() {
                user := u.Value.(*OnlineUser)
                user.Send <- bc
            }
        }
    }
}
