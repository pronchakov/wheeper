package main

import (
	"github.com/pronchakov/wheeper/pkg/dao/moonraker"
	"github.com/pronchakov/wheeper/pkg/logic"
	"log"
	"time"
)

func main() {

	moonraker.SetUrl("ws://ray.local:7125/websocket")
	moonraker.Connect()
	moonraker.Subscribe()

	statusUpdateChannel := make(chan logic.StatusUpdate)
	go moonraker.HandleMessages(statusUpdateChannel)
	go logic.ConstantlyUpdateMainData(statusUpdateChannel)

	masterData := logic.GetMasterData()

	for i := 0; i < 60; i++ {
		if masterData.Objects.Extruder != nil && masterData.Objects.Extruder.Temperature != nil {
			log.Println("Data:", *masterData.Objects.Extruder.Temperature)
		} else {
			log.Println("not yet")
		}
		time.Sleep(time.Second)
	}

	moonraker.Disconnect()
}
