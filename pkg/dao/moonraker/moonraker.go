package moonraker

import (
	"context"
	"encoding/json"
	"github.com/pronchakov/wheeper/pkg/logic"
	"log"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

var url string
var username string
var password string

var ctx context.Context
var cancel context.CancelFunc
var conn *websocket.Conn
var err error

func SetUrl(u string) {
	url = u
}

func SetUsername(u string) {
	username = u
}

func SetPassword(p string) {
	password = p
}

func Connect() {
	ctx, cancel = context.WithTimeout(context.Background(), time.Minute)
	conn, _, err = websocket.Dial(ctx, url, nil)
	if err != nil {
		log.Fatal("Error:", err) // todo:
	}
}

func HandleMessages(outputChannel chan<- logic.StatusUpdate) {
	for {
		_, readBytes, readError := conn.Read(ctx)
		if readError != nil {
			log.Fatal("Error processing message:", readError)
		}

		m := map[string]string{}
		json.Unmarshal(readBytes, &m)

		method := m["method"]
		if method == "notify_status_update" {
			statusUpdate := logic.StatusUpdate{}
			json.Unmarshal(readBytes, &statusUpdate)
			outputChannel <- statusUpdate
		}
	}
}

func Subscribe() {
	request := logic.NewSubscribeRequest()
	err2 := wsjson.Write(ctx, conn, request)
	if err2 != nil {
		log.Fatal("Error subscribing:", err2)
	}
}

func Disconnect() {
	conn.Close(websocket.StatusNormalClosure, "end")
	conn.CloseNow()
	cancel()
}
