package middleware

import (
	config "NGB-MSG-handler/internal/conf"
	"NGB-MSG-handler/internal/model"
	"NGB-MSG-handler/internal/util"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/net/websocket"
)

var receiverChannel *amqp.Channel
var msgs <-chan amqp.Delivery

func RabbitMQInit() {
	//establish connection
	dialStr := fmt.Sprintf("amqp://%s:%s@%s:%s", config.Config.AMQPConfig.User,
		config.Config.AMQPConfig.Password,
		config.Config.AMQPConfig.Host,
		config.Config.AMQPConfig.Port)
	conn, err := amqp.Dial(dialStr)
	if err != nil {
		util.MakeInfoLog("Failed to connect to RabbitMQ")
	}
	//create channel
	receiverChannel, err = conn.Channel()
	if err != nil {
		util.MakeInfoLog("[RabbitMQ]Failed to open a channel")
	}
	//decleare queue
	queue, err := receiverChannel.QueueDeclare(
		"userActivity",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		util.MakeErrorLog("[RabbitMQ]Failed to decleare a queue")

	} else {
		util.MakeInfoLog(fmt.Sprintf("[RabbitMQ]Queue %s Decleared", queue.Name))
	}
	ReceiveByteFromQueue()
}

func ReceiveByteFromQueue() {
	var err error
	msgs, err = receiverChannel.Consume(
		"userActivity",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		util.MakeInfoLog("[RabbitMQ]Failed to consume a message")
	} else {
		util.MakeInfoLog("[RabbitMQ]Start to receive message")
	}
	go messageStoragePersistent()
}

func messageStoragePersistent() {
	util.MakeInfoLog("[Storage]Starting to store message into database")
	var msgUnmarshaled model.Message
	for msgIterator := range msgs {
		err := json.Unmarshal(msgIterator.Body, &msgUnmarshaled)
		if err != nil {
			util.MakeErrorLog("[storage]failed to unmarshal a message")
			continue
		}
		model.CreateMessage(msgUnmarshaled.ContentType, msgUnmarshaled.Body, msgUnmarshaled.TargetUid)
		util.MakeInfoLog("[Storage]message saved")
	}

}

//
//Run in a goroutine
//
func PushingMessageToClient(ws *websocket.Conn, clientUid string) {
	util.MakeInfoLog("pushing message")
	msgs, err := model.QueryMessageByUid(clientUid)
	if err != nil {
		util.MakeErrorLog("[model]" + err.Error())
		return
	}
	msgMarshaled, err := json.Marshal(&msgs)
	if err != nil {
		util.MakeErrorLog("[json marshal]" + err.Error())
	}
	ws.Write(msgMarshaled)

}
