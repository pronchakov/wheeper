package moonraker

import (
	"context"
	"dario.cat/mergo"
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

var needToHandleMessages bool = true

func ConnectToMoonraker() {
	ctx, cancel = context.WithTimeout(context.Background(), time.Minute)
	conn, _, err = websocket.Dial(ctx, "ws://ray.local:7125/websocket", nil)
	if err != nil {
		log.Fatal("Error:", err) // todo:
	}
}

func HandleMessages(data *logic.Data) {
	for needToHandleMessages == true {
		_, readBytes, readError := conn.Read(ctx)
		if readError != nil {
			log.Fatal("Error processing message:", readError)
		}

		m := map[string]any{}
		json.Unmarshal(readBytes, &m)

		method := m["method"]
		if method == "notify_status_update" {
			statusUpdate := logic.StatusUpdate{}
			json.Unmarshal(readBytes, &statusUpdate)
			for _, p := range statusUpdate.Params {
				if p.Extruder != nil {
					mergo.Merge(data.Objects.Extruder, p.Extruder, mergo.WithOverride)
				}
			}
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

func Close() {
	needToHandleMessages = false
	time.Sleep(time.Second)
	conn.Close(websocket.StatusNormalClosure, "end")
	conn.CloseNow()
	cancel()
}
