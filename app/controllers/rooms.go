package controllers

import "github.com/robfig/revel"

type Rooms struct {
	*revel.Controller
}

func (c Rooms) Index() revel.Result {
	return c.Render()
}
