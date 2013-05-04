package chatserver 

import (
    //"code.google.com/p/go.net/websocket"
    "webchat/app/model"
    "container/list" 
    "fmt"
)

type Server struct {
    Name    string
    ActiveRooms *list.List
}

type Event struct {
    Type string
    Text string
    User *UserInfo
}

func NewServer() *Server {
    Fx := &Server{
        Name: "webchat",
        ActiveRooms: list.New(),
    }
    return Fx
}

// find avtive room return a activeroom instance
func(s *Server) GetActiveRoom(roomkey string) *ActiveRoom {
    var activeroom *ActiveRoom
    for room := s.ActiveRooms.Front(); room != nil; room = room.Next() {
        r := room.Value.(*ActiveRoom)
        if r.RoomKey == roomkey {
            activeroom = r
        }
    }

    if activeroom == nil {
        activeroom = NewActiveRoom(roomkey)
        go activeroom.Run()
        s.ActiveRooms.PushBack(activeroom)
    }

    return activeroom
}

// init all room 
func (s *Server) RunRooms() {
    rooms := model.AllRoom()

    for _, room := range rooms {
        fmt.Println(room)
        activeroom := NewActiveRoom(room.RoomKey)
        go activeroom.Run()
        s.ActiveRooms.PushBack(activeroom)
    }
}
