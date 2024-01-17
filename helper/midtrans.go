package helper

import (
	"be-project/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func InitMidtrans() (snap.Client, midtrans.EnvironmentType, string) {
	var snap snap.Client
	conf := config.GetConfig()
	if conf.Midtrans.EnvironmentType == "Sandbox" {
		midtrans.Environment = midtrans.Sandbox
		midtrans.ServerKey = conf.Midtrans.Sandbox.Server
		snap.Env = midtrans.Environment
		snap.ServerKey = conf.Midtrans.Sandbox.Server
	} else {
		midtrans.Environment = midtrans.Production
		midtrans.ServerKey = conf.Midtrans.Production.Server
		snap.Env = midtrans.Environment
		snap.ServerKey = conf.Midtrans.Production.Server
	}

	snap.New(midtrans.ServerKey, midtrans.Environment)
	return snap, midtrans.Environment, midtrans.ServerKey

}
