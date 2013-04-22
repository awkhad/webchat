package controllers


import (
    "github.com/robfig/revel"
    "webchat/app/form"
    "webchat/app/model"
    //"fmt"
)


type Rooms struct {
    *Application
}

func (c Rooms) Index() revel.Result {
	return c.Render()
}

func (c Rooms) New() revel.Result {

    if !c.isLogin() {
        c.Flash.Error("Please login first")
        return c.Redirect(Application.Index)
    }

	return c.Render()
}

func (c Rooms) Create(rf *form.RoomForm) revel.Result {

    if !c.isLogin() {
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

    return c.Redirect("/rooms/%s/show", room.RoomKey)
}

func (c Rooms) Show(roomkey string) revel.Result{
    return c.Render(roomkey)
}
