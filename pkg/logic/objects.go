package logic

type Webhooks struct {
	State        *string `json:"state"`
	StateMessage *string `json:"state_message"`
}

type GcodeMove struct {
	SpeedFactor         *float32   `json:"speed_factor"`
	Speed               *float32   `json:"speed"`
	ExtrudeFactor       *float32   `json:"extrude_factor"`
	AbsoluteCoordinates *bool      `json:"absolute_coordinates"`
	AbsoluteExtrude     *bool      `json:"absolute_extrude"`
	HomingOrigin        *[]float32 `json:"homing_origin"`
	Position            *[]float32 `json:"position"`
	GcodePosition       *[]float32 `json:"gcode_position"`
}

type Toolhead struct {
	HomedAxes            *string    `json:"homed_axes"`
	PrintTime            *float32   `json:"print_time"`
	EstimatedPrintTime   *float32   `json:"estimated_print_time"`
	Extruder             *string    `json:"extruder"`
	Position             *[]float32 `json:"position"`
	MaxVelocity          *float32   `json:"max_velocity"`
	MaxAccel             *float32   `json:"max_accel"`
	MaxAccelToDecel      *float32   `json:"max_accel_to_decel"`
	SquareCornerVelocity *float32   `json:"square_corner_velocity"`
}

type Extruder struct {
	Temperature     *float32 `json:"temperature"`
	Target          *float32 `json:"target"`
	Power           *float32 `json:"power"`
	PressureAdvance *float32 `json:"pressure_advance"`
	SmoothTime      *float32 `json:"smooth_time"`
}

type HeaterBed struct {
	Temperature *float32 `json:"temperature"`
	Target      *float32 `json:"target"`
	Power       *float32 `json:"power"`
}

type Fan struct {
	Speed *float32 `json:"speed"`
	Rpm   *int     `json:"rpm"`
}

type IdleTimeout struct {
	State        *string  `json:"state"`
	PrintingTime *float32 `json:"printing_time"`
}

type PrintStats struct {
	Filename      *string         `json:"filename"`
	TotalDuration *float32        `json:"total_duration"`
	PrintDuration *float32        `json:"print_duration"`
	FilamentUsed  *float32        `json:"filament_used"`
	State         *string         `json:"state"`
	Message       *string         `json:"message"`
	Info          *PrintStatsInfo `json:"info"`
}

type PrintStatsInfo struct {
	TotalLayer   *interface{} `json:"total_layer"`
	CurrentLayer *interface{} `json:"current_layer"`
}
