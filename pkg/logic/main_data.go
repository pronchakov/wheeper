package logic

type Data struct {
	Objects *PrinterObjects
}

type PrinterObjects struct {
	Webhooks    *Webhooks
	GcodeMove   *GcodeMove
	Toolhead    *Toolhead
	Extruder    *Extruder
	HeaterBed   *HeaterBed
	Fan         *Fan
	IdleTimeout *IdleTimeout
	PrintStats  *PrintStats
}

func NewData() *Data {
	return &Data{Objects: &PrinterObjects{
		Webhooks:    &Webhooks{},
		GcodeMove:   &GcodeMove{},
		Toolhead:    &Toolhead{},
		Extruder:    &Extruder{},
		HeaterBed:   &HeaterBed{},
		Fan:         &Fan{},
		IdleTimeout: &IdleTimeout{},
		PrintStats:  &PrintStats{Info: &PrintStatsInfo{}},
	}}
}
