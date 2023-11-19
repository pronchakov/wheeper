package main

import (
	"github.com/pronchakov/wheeper/pkg/dao/moonraker"
	"github.com/pronchakov/wheeper/pkg/logic"
	"log"
	"time"
)

func main() {

	data := logic.NewData()

	moonraker.ConnectToMoonraker()
	moonraker.Subscribe()
	go moonraker.HandleMessages(data)

	for i := 0; i < 60; i++ {
		if data.Objects.Extruder != nil && data.Objects.Extruder.Temperature != nil {
			log.Println("Data:", *data.Objects.Extruder.Temperature)
		}
		time.Sleep(time.Second)
	}

	time.Sleep(time.Hour)
}
