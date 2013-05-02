package chatserver 

import (
    "container/list" 
    "fmt"
)

func init(){
    server := &Server{
        Name: "webchat",
        Rooms: list.New(),
    }

    fmt.Println("------server run...------ server name is", server.Name)
}

type Server struct {
    Name string
    Rooms *list.List
}
