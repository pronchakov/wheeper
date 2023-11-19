package main

import (
	"github.com/pronchakov/wheeper/pkg/dao/moonraker"
	"github.com/pronchakov/wheeper/pkg/logic"
	"log"
	"time"
)

func main() {

	masterData := logic.NewData()

	moonraker.SetUrl("ws://ray.local:7125/websocket")
	moonraker.Connect()
	moonraker.Subscribe()
	go moonraker.HandleMessages(masterData)

	for i := 0; i < 60; i++ {
		if masterData.Objects.Extruder != nil && masterData.Objects.Extruder.Temperature != nil {
			log.Println("Data:", *masterData.Objects.Extruder.Temperature)
		}
		time.Sleep(time.Second)
	}

	moonraker.Disconnect()
}
