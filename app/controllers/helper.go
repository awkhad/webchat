package controllers

import (
    "github.com/robfig/revel"
    "webchat/app/model"
    //"fmt"
)

func (c Application) isLogin() bool {
    if _, ok := c.Session["user_name"]; ok {
        return true
    }
    return false
}

func (c Application) CheckUser() revel.Result {
    if !c.isLogin() {
        c.Flash.Error("Please login first")
        //fmt.Println(c.Flash)
        return c.Redirect(Application.Index)
    }
    return nil
}

func (c Application) AddUser() revel.Result {
    //fmt.Println(c.Session["user_name"])
    if c.isLogin() {
        user := model.FindUserByName(c.Session["user_name"])
        c.RenderArgs["user"] = user 
    }
    return nil
}

func (c Application) CurrentUser() (user *model.User) {
    if c.isLogin() {
        user = model.FindUserByName(c.Session["user_name"])
    }else {
        return nil
    }
    return user
}
