package controllers


import (
    "github.com/robfig/revel"
    //"fmt"
)


type Rooms struct {
    *Application
	//*revel.Controller
}

func (c Rooms) Index() revel.Result {
    //fmt.Println(c.isLogin())
	return c.Render()
}

func (c Rooms) New() revel.Result {
    if !c.isLogin() {
        c.Flash.Error("Please login first")
        return c.Redirect(Application.Index)
    }
	return c.Render()
}
