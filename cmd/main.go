package main

import (
	config "NGB-MSG-handler/internal/conf"
	"NGB-MSG-handler/internal/util"
)

func main() {
	config.ReadConfig()
	util.LogUtilInit()
	util.MakeInfoLog("service started")
}
