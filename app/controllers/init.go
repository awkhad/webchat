package controllers

import (
    "github.com/robfig/revel"
)

func init(){
    //revel.InterceptMethod(Rooms.CheckUser, revel.BEFORE)
    revel.InterceptMethod(Application.AddUser, revel.BEFORE)
}
