package moonraker

import (
	"context"
	"github.com/google/uuid"
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

	go handleMessages()

}

func handleMessages() {
	for needToHandleMessages == true {
		v := map[string]any{}
		err := wsjson.Read(ctx, conn, &v)
		if err != nil {
			log.Println("Error processing message:", err, v)
		}
		method := v["method"]
		if method == "notify_proc_stat_update" {
			log.Println("Message:", method)
		}
		if method == "notify_status_update" {
			log.Println("Message:", method, v["params"])
		}
	}
}

func Subscribe(object string) {
	request := SubscribeRequest{
		Jsonrpc: "2.0",
		Method:  "printer.objects.subscribe",
		Params: Params{
			Objects: Objects{
				Extruder: nil,
			},
		},
		Id: uuid.New().String(),
	}
	err2 := wsjson.Write(ctx, conn, &request)
	if err2 != nil {
		log.Println("Error subscribing:", err2)
	}
}

func Close() {
	needToHandleMessages = false
	time.Sleep(time.Second)
	conn.Close(websocket.StatusNormalClosure, "end")
	conn.CloseNow()
	cancel()
}
