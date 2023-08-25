package controller

import (
	"NGB-MSG-handler/internal/middleware"
	"NGB-MSG-handler/internal/util"
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

type ConnInfoStruct struct {
	ClientUid string `json:"clientUid"`
}

// Echo the data received on the WebSocket.
func FetchMessage(ws *websocket.Conn) {
	defer ws.Close()
	msgMarshaled := make([]byte, 512)
	var msgUnmarshaled ConnInfoStruct
	for {
		n, err := ws.Read(msgMarshaled)
		if err != nil {
			util.MakeErrorLog("[FetchMessage]" + err.Error())
			return
		}
		json.Unmarshal(msgMarshaled[:n], &msgUnmarshaled)
		fmt.Println(string(msgUnmarshaled.ClientUid))
		middleware.PushingMessageToClient(ws, msgUnmarshaled.ClientUid)
		time.Sleep(time.Duration(5) * time.Second)
	}

}
