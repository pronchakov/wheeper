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

func HandleMessages(data *logic.Data) {
	for needToHandleMessages == true {
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
			for _, p := range statusUpdate.Params {
				if p.Webhooks != nil {
					mergo.Merge(data.Objects.Webhooks, p.Webhooks, mergo.WithOverride)
				}
				if p.GcodeMove != nil {
					mergo.Merge(data.Objects.GcodeMove, p.GcodeMove, mergo.WithOverride)
				}
				if p.Toolhead != nil {
					mergo.Merge(data.Objects.Toolhead, p.Toolhead, mergo.WithOverride)
				}
				if p.Extruder != nil {
					mergo.Merge(data.Objects.Extruder, p.Extruder, mergo.WithOverride)
				}
				if p.HeaterBed != nil {
					mergo.Merge(data.Objects.HeaterBed, p.HeaterBed, mergo.WithOverride)
				}
				if p.Fan != nil {
					mergo.Merge(data.Objects.Fan, p.Fan, mergo.WithOverride)
				}
				if p.IdleTimeout != nil {
					mergo.Merge(data.Objects.IdleTimeout, p.IdleTimeout, mergo.WithOverride)
				}
				if p.PrintStats != nil {
					mergo.Merge(data.Objects.PrintStats, p.PrintStats, mergo.WithOverride)
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

func Disconnect() {
	needToHandleMessages = false
	time.Sleep(time.Second)
	conn.Close(websocket.StatusNormalClosure, "end")
	conn.CloseNow()
	cancel()
}
