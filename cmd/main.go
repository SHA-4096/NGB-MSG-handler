package main

import (
	config "NGB-MSG-handler/internal/conf"
	"NGB-MSG-handler/internal/middleware"
	"NGB-MSG-handler/internal/model"
	"NGB-MSG-handler/internal/router"
	"NGB-MSG-handler/internal/util"
	"net/http"
)

func main() {
	config.ReadConfig()
	util.LogUtilInit()
	model.MySqlInit()
	go middleware.RabbitMQInit()
	router.RouterInit()
	WsServerInit()
	util.MakeInfoLog("service started")
	var forever chan struct{}
	<-forever
}

// This example demonstrates a trivial echo server.
func WsServerInit() {
	util.MakeInfoLog("Serving websocket")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
