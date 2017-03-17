// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApplication struct {}
var Application tApplication


func (_ tApplication) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Index", args).Url
}

func (_ tApplication) EnterDemo(
		user string,
		demo string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "demo", demo)
	return revel.MainRouter.Reverse("Application.EnterDemo", args).Url
}


type tUsers struct {}
var Users tUsers


func (_ tUsers) Login(
		name string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Users.Login", args).Url
}

func (_ tUsers) Register(
		name string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Users.Register", args).Url
}

func (_ tUsers) State(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.State", args).Url
}

func (_ tUsers) Online(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.Online", args).Url
}

func (_ tUsers) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.Logout", args).Url
}


type tRefresh struct {}
var Refresh tRefresh


func (_ tRefresh) Index(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Refresh.Index", args).Url
}

func (_ tRefresh) Room(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Refresh.Room", args).Url
}

func (_ tRefresh) Say(
		user string,
		message string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "message", message)
	return revel.MainRouter.Reverse("Refresh.Say", args).Url
}

func (_ tRefresh) Leave(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Refresh.Leave", args).Url
}


type tWebSocket struct {}
var WebSocket tWebSocket


func (_ tWebSocket) Room(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("WebSocket.Room", args).Url
}

func (_ tWebSocket) RoomSocket(
		user string,
		ws interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "ws", ws)
	return revel.MainRouter.Reverse("WebSocket.RoomSocket", args).Url
}


type tLongPolling struct {}
var LongPolling tLongPolling


func (_ tLongPolling) Room(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("LongPolling.Room", args).Url
}

func (_ tLongPolling) Say(
		user string,
		message string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "message", message)
	return revel.MainRouter.Reverse("LongPolling.Say", args).Url
}

func (_ tLongPolling) PrevPage(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("LongPolling.PrevPage", args).Url
}

func (_ tLongPolling) WaitMessages(
		last uint64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "last", last)
	return revel.MainRouter.Reverse("LongPolling.WaitMessages", args).Url
}

func (_ tLongPolling) LoadPage(
		last uint64,
		page uint64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "last", last)
	revel.Unbind(args, "page", page)
	return revel.MainRouter.Reverse("LongPolling.LoadPage", args).Url
}

func (_ tLongPolling) Leave(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("LongPolling.Leave", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
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


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


