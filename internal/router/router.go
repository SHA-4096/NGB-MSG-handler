package router

import (
	"NGB-MSG-handler/internal/controller"
	"net/http"

	"golang.org/x/net/websocket"
)

func RouterInit() {
	http.Handle("/fetch", websocket.Handler(controller.FetchMessage))
}
