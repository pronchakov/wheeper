package logic

import "github.com/google/uuid"

type Objects struct {
	Webhooks    *[]string `json:"webhooks"`
	GcodeMove   *[]string `json:"gcode_move"`
	Toolhead    *[]string `json:"toolhead"`
	Extruder    *[]string `json:"extruder"`
	HeaterBed   *[]string `json:"heater_bed"`
	Fan         *[]string `json:"fan"`
	IdleTimeout *[]string `json:"idle_timeout"`
	PrintStats  *[]string `json:"print_stats"`
}

type Params struct {
	Objects Objects `json:"objects"`
}

type SubscribeRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
	Id      string `json:"id"`
}

func NewSubscribeRequest() *SubscribeRequest {
	return &SubscribeRequest{
		Jsonrpc: "2.0",
		Method:  "printer.objects.subscribe",
		Params: Params{Objects: Objects{
			Webhooks:    nil,
			GcodeMove:   nil,
			Toolhead:    nil,
			Extruder:    nil,
			HeaterBed:   nil,
			Fan:         nil,
			IdleTimeout: nil,
			PrintStats:  nil,
		}},
		Id: uuid.New().String(),
	}
}
