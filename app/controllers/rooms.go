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
