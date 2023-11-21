package logic

import (
	"dario.cat/mergo"
	"log"
)

var masterData = NewData()

func GetMasterData() *Data {
	return masterData
}

func ConstantlyUpdateMainData(inputChannel <-chan StatusUpdate) {
	for {
		statusUpdate := <-inputChannel
		log.Println("Updating master data")

		for _, p := range statusUpdate.Params {
			if p.Webhooks != nil {
				mergo.Merge(masterData.Objects.Webhooks, p.Webhooks, mergo.WithOverride)
			}
			if p.GcodeMove != nil {
				mergo.Merge(masterData.Objects.GcodeMove, p.GcodeMove, mergo.WithOverride)
			}
			if p.Toolhead != nil {
				mergo.Merge(masterData.Objects.Toolhead, p.Toolhead, mergo.WithOverride)
			}
			if p.Extruder != nil {
				mergo.Merge(masterData.Objects.Extruder, p.Extruder, mergo.WithOverride)
			}
			if p.HeaterBed != nil {
				mergo.Merge(masterData.Objects.HeaterBed, p.HeaterBed, mergo.WithOverride)
			}
			if p.Fan != nil {
				mergo.Merge(masterData.Objects.Fan, p.Fan, mergo.WithOverride)
			}
			if p.IdleTimeout != nil {
				mergo.Merge(masterData.Objects.IdleTimeout, p.IdleTimeout, mergo.WithOverride)
			}
			if p.PrintStats != nil {
				mergo.Merge(masterData.Objects.PrintStats, p.PrintStats, mergo.WithOverride)
			}
		}
	}
}
