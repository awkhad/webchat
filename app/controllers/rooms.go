package controllers


import (
    "github.com/robfig/revel"
    "webchat/app/form"
    "webchat/app/model"
    "fmt"
)


type Rooms struct {
    *Application
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

    rf.UserId = c.CurrentUser().Id

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

    return c.Redirect("/r/%s", room.RoomKey)
}

func (c Rooms) Show(roomkey string) revel.Result{
    if !isLogin(c.Controller) {
        c.Flash.Error("Please login first")
        return c.Redirect(Application.Index)
    }

    room := model.FindRoomByRoomKey(roomkey)

    activeRoom := ChatServer.GetActiveRoom(roomkey)
    users := activeRoom.UserList()

    return c.Render(room, users)
}
