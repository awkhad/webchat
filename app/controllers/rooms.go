package controllers

import (
	"fmt"
	"github.com/robfig/revel"
	"webchat/app/chatserver"
	"webchat/app/form"
	"webchat/app/model"
)

type Rooms struct {
	*Application
}

type RoomApi struct {
	*revel.Controller
}

type UserList struct {
	Users []*chatserver.UserInfo
}

func (c Rooms) Index(p int) revel.Result {
	fmt.Println("p is:", p)
	if p == 0 {
		p = 1
	}
	rooms := model.FindOnePage(p)
	allPage := (model.RoomCount() + model.PageSize - 1) / model.PageSize
	return c.Render(rooms, p, allPage)
}

func (c Rooms) New() revel.Result {

	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	return c.Render()
}

func (c Rooms) Create(rf *form.RoomForm) revel.Result {

	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	rf.UserId = CurrentUser(c.Controller).Id

	rf.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Rooms.New)
	}
	room := model.NewRoom(rf)

	if _, err := room.Save(); err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(Rooms.New)
	}

	// run activeroom
	activeroom := chatserver.NewActiveRoom(room.RoomKey)
	go activeroom.Run()
	ChatServer.ActiveRooms.PushBack(activeroom)

	return c.Redirect("/r/%s", room.RoomKey)
}

func (c Rooms) Show(roomkey string) revel.Result {
	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	room := model.FindRoomByRoomKey(roomkey)

	activeRoom := ChatServer.GetActiveRoom(roomkey)
	users := activeRoom.UserList()

	return c.Render(room, users)
}

func (c Rooms) Edit(roomkey string) revel.Result {

	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	room := model.FindRoomByRoomKey(roomkey)

	return c.Render(room)
}

func (c Rooms) Update(roomkey string, updateroom *form.UpdateRoom) revel.Result {

	if !isLogin(c.Controller) {
		c.Flash.Error("Please login first")
		return c.Redirect(Application.Index)
	}

	room := model.FindRoomByRoomKey(roomkey)

	if err := room.Update(updateroom); err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect("/r/%s/edit", room.RoomKey)
	}

	c.Flash.Success("update success")
	return c.Redirect("/r/%s/edit", room.RoomKey)
}

func (c RoomApi) Users(roomkey string) revel.Result {

	// get a activeRoom and get room's user list 
	activeroom := ChatServer.GetActiveRoom(roomkey)
	users := activeroom.UserList()

	userList := &UserList{
		Users: users,
	}

	return c.RenderJson(userList)
}
