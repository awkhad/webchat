package chatserver 

import (
    "container/list" 
    "fmt"
)

type Server struct {
    Name    string
    Rooms   *list.List
}

type Room struct {
    Id      int 
    RoomKey string
    Users   *list.List
}

type User struct {
    Id int
}

type Event struct {
    Type string
    Text string
}

func NewServer() *Server {
    Fx := &Server{
        Name: "webchat",
        Rooms: list.New(),
    }
    return Fx
}

func (r Room)JoinUser(user string) {
    fmt.Println("user is:", user)
}



