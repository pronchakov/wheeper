package moonraker

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

type Webhooks struct {
	State        *string `json:"state"`
	StateMessage *string `json:"state_message"`
}

type GcodeMove struct {
	SpeedFactor         *float64   `json:"speed_factor"`
	Speed               *float64   `json:"speed"`
	ExtrudeFactor       *float64   `json:"extrude_factor"`
	AbsoluteCoordinates *bool      `json:"absolute_coordinates"`
	AbsoluteExtrude     *bool      `json:"absolute_extrude"`
	HomingOrigin        *[]float64 `json:"homing_origin"`
	Position            *[]float64 `json:"position"`
	GcodePosition       *[]float64 `json:"gcode_position"`
}

type Toolhead struct {
	HomedAxes            *string    `json:"homed_axes"`
	PrintTime            *float64   `json:"print_time"`
	EstimatedPrintTime   *float64   `json:"estimated_print_time"`
	Extruder             *string    `json:"extruder"`
	Position             *[]float64 `json:"position"`
	MaxVelocity          *float64   `json:"max_velocity"`
	MaxAccel             *float64   `json:"max_accel"`
	MaxAccelToDecel      *float64   `json:"max_accel_to_decel"`
	SquareCornerVelocity *float64   `json:"square_corner_velocity"`
}

type Extruder struct {
	Temperature     *float32 `json:"temperature"`
	Target          *float32 `json:"target"`
	Power           *float32 `json:"power"`
	PressureAdvance *float32 `json:"pressure_advance"`
	SmoothTime      *float32 `json:"smooth_time"`
}

type HeaterBed struct {
	Temperature *float64 `json:"temperature"`
	Target      *float64 `json:"target"`
	Power       *float64 `json:"power"`
}

type Fan struct {
	Speed *float64 `json:"speed"`
	Rpm   *int     `json:"rpm"`
}

type IdleTimeout struct {
	State        *string  `json:"state"`
	PrintingTime *float64 `json:"printing_time"`
}

type PrintStats struct {
	Filename      *string  `json:"filename"`
	TotalDuration *float64 `json:"total_duration"`
	PrintDuration *float64 `json:"print_duration"`
	FilamentUsed  *float64 `json:"filament_used"`
	State         *string  `json:"state"`
	Message       *string  `json:"message"`
	Info          *Info    `json:"info"`
}

type Info struct {
	TotalLayer   *interface{} `json:"total_layer"`
	CurrentLayer *interface{} `json:"current_layer"`
}
