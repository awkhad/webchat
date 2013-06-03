package chatserver

import (
	"code.google.com/p/go.net/websocket"
	"webchat/app/model"
	//"strconv"
	"fmt"
	"time"
)

type OnlineUser struct {
	Id         int
	Connection *websocket.Conn
	Send       chan *Event
	Room       *ActiveRoom
	Info       *UserInfo
}

type UserInfo struct {
	Name   string
	Email  string
	Avatar string
}

func NewOnlineUser(user *model.User, ws *websocket.Conn, room *ActiveRoom) *OnlineUser {
	onlineUser := &OnlineUser{
		Id:         user.Id,
		Connection: ws,
		Send:       make(chan *Event, 512),
		Room:       room,

		Info: &UserInfo{
			Name:   user.Name,
			Email:  user.Email,
			Avatar: user.AvatarUrl(),
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
		event.Created = time.Now()
		event.User = u.Info
		fmt.Println("the message is:", event)

		// user close
		if err != nil {
			fmt.Println("Receive occur some error", err.Error())
			return
		}

		u.Room.Broadcast <- &event
		u.SaveMessageToRedis(&event)
	}
}

func (u *OnlineUser) SaveMessageToRedis(event *Event) {
	// save to redis list
	// format: "text|lds|asd"
	listKey := "room:" + u.Room.RoomKey
	//time := event.Created.Unix()
	content := event.Type + "|" + event.User.Name + "|" + event.Text + "|" + event.Created.String()
	redisClient.Lpush(listKey, []byte(content))
}

func (u *OnlineUser) Close() {
	// clear resource when user conn close
	// close conn
	if err := u.Connection.Close(); err != nil {
		fmt.Println("close conn faild")
	}

	// close channel
	close(u.Send)

	// send levae message to other client
	event := &Event{
		Type:    "leave",
		Text:    u.Info.Name + " has leave room",
		User:    u.Info,
		Created: time.Now(),
	}

	u.Room.Broadcast <- event
}
