package logic

type StatusUpdate struct {
	Jsonrpc string  `json:"jsonrpc"`
	Method  string  `json:"method"`
	Params  []Param `json:"params"`
}

type Param struct {
	Webhooks    *Webhooks    `json:"webhooks"`
	GcodeMove   *GcodeMove   `json:"gcode_move"`
	Toolhead    *Toolhead    `json:"toolhead"`
	Extruder    *Extruder    `json:"extruder"`
	HeaterBed   *HeaterBed   `json:"heater_bed"`
	Fan         *Fan         `json:"fan"`
	IdleTimeout *IdleTimeout `json:"idle_timeout"`
	PrintStats  *PrintStats  `json:"print_stats"`
}
