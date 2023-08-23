package main

import (
	config "NGB-MSG-handler/internal/conf"
	"NGB-MSG-handler/internal/middleware"
	"NGB-MSG-handler/internal/model"
	"NGB-MSG-handler/internal/util"
)

func main() {
	config.ReadConfig()
	util.LogUtilInit()
	model.MySqlInit()
	middleware.RabbitMQInit()
	util.MakeInfoLog("service started")
	var forever chan struct{}
	<-forever
}
