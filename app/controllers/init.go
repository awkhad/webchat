package controllers

import (
    "github.com/robfig/revel"
)

func init(){
    //revel.InterceptMethod(Rooms.CheckUser, revel.BEFORE)
    revel.InterceptMethod(Application.AddUser, revel.BEFORE)

    revel.TemplateFuncs["ueq"] = func(a, b interface{}) bool { return !(a == b) }
    revel.TemplateFuncs["add"] = func(a, b int) int { return a + b }
    revel.TemplateFuncs["minus"] = func(a, b int) int { return a - b }
    revel.TemplateFuncs["less"] = func(a, b int) bool { return a < b }
}
