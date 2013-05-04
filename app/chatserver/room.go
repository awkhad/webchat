package chatserver 

import (
    //"code.google.com/p/go.net/websocket"
    "container/list" 
)

type ActiveRoom struct {
    RoomKey string
    Users   *list.List
    Broadcast chan Event
}

func NewActiveRoom(rk string) *ActiveRoom{
    activeRoom := &ActiveRoom{
        RoomKey: rk,
        Users: list.New(),
        Broadcast: make(chan Event),
    }
    return activeRoom
}

func (r ActiveRoom) JoinUser(user *OnlineUser) {
    r.Users.PushBack(user)
}

func (r ActiveRoom) UserList() ([]*OnlineUser){
    var userList []*OnlineUser

    for u := r.Users.Front(); u != nil; u = u.Next() {
        user := u.Value.(*OnlineUser)
        userList = append(userList, user)
    }
    return userList
}

func (r ActiveRoom) run() {
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
