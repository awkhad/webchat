package chatserver

import (
	"code.google.com/p/go.net/websocket"
	"webchat/app/model"
	//"container/list" 
	"fmt"
)

type OnlineUser struct {
	Id         int
	Connection *websocket.Conn
	Send       chan *Event
	Room       *ActiveRoom
	Info       *UserInfo
}

type UserInfo struct {
	Name  string
	Email string
}

func NewOnlineUser(user *model.User, ws *websocket.Conn, room *ActiveRoom) *OnlineUser {
	onlineUser := &OnlineUser{
		Id:         user.Id,
		Connection: ws,
		Send:       make(chan *Event, 512),
		Room:       room,
		Info: &UserInfo{
			Name:  user.Name,
			Email: user.Email,
		},
	}
	return onlineUser
}

func (u *OnlineUser) PushToClient() {
	for b := range u.Send {
		err := websocket.JSON.Send(u.Connection, b)
		if err != nil {
			break
		}
	}
}

func (u *OnlineUser) PullFromClient() {
	for {
		var event Event
		err := websocket.JSON.Receive(u.Connection, &event)
        event.User = u.Info
		fmt.Println("the message is:", event)

		// user close
		if err != nil {
			fmt.Println("Receive occur some error", err.Error())
			return
		}
		u.Room.Broadcast <- &event
	}
}

func (u *OnlineUser) Close() {
	// clear resource when user conn close 
	// remove user form rooms's users list 
	for e := u.Room.Users.Front(); e != nil; e = e.Next() {
		user := e.Value.(*OnlineUser)
		if user.Id == u.Id && user.Connection == u.Connection {
			u.Room.Users.Remove(e)
			break
		}
	}
	// close conn
	if err := u.Connection.Close(); err != nil {
		fmt.Println("close conn faild")
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
