// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/robfig/revel"


type tWebsocket struct {}
var Websocket tWebsocket


func (p tWebsocket) Chat(
		roomkey string,
		ws interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomkey", roomkey)
	revel.Unbind(args, "ws", ws)
	return revel.MainRouter.Reverse("Websocket.Chat", args).Url
}


type tRoomApi struct {}
var RoomApi tRoomApi


func (p tRoomApi) Users(
		roomkey string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomkey", roomkey)
	return revel.MainRouter.Reverse("RoomApi.Users", args).Url
}


type tApplication struct {}
var Application tApplication


func (p tApplication) CheckUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.CheckUser", args).Url
}

func (p tApplication) AddUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.AddUser", args).Url
}

func (p tApplication) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Index", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (p tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (p tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (p tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (p tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (p tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tSessions struct {}
var Sessions tSessions


func (p tSessions) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sessions.New", args).Url
}

func (p tSessions) Create(
		loginform interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "loginform", loginform)
	return revel.MainRouter.Reverse("Sessions.Create", args).Url
}

func (p tSessions) Destroy(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sessions.Destroy", args).Url
}


type tRooms struct {}
var Rooms tRooms


func (p tRooms) Index(
		p int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "p", p)
	return revel.MainRouter.Reverse("Rooms.Index", args).Url
}

func (p tRooms) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Rooms.New", args).Url
}

func (p tRooms) Create(
		rf interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "rf", rf)
	return revel.MainRouter.Reverse("Rooms.Create", args).Url
}

func (p tRooms) Show(
		roomkey string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roomkey", roomkey)
	return revel.MainRouter.Reverse("Rooms.Show", args).Url
}


type tUsers struct {}
var Users tUsers


func (p tUsers) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.New", args).Url
}

func (p tUsers) Create(
		userform interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userform", userform)
	return revel.MainRouter.Reverse("Users.Create", args).Url
}

func (p tUsers) EditSettings(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.EditSettings", args).Url
}

func (p tUsers) SaveSettings(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.SaveSettings", args).Url
}

func (p tUsers) MyRooms(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.MyRooms", args).Url
}


