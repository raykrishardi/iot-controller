package main

import (
	"context"

	"github.com/raykrishardi/iot-controller-grpc/internal/pkg/iot"
	"github.com/raykrishardi/iot-controller-grpc/internal/pkg/led"
	"google.golang.org/grpc/status"
)

func (i *IotServer) ProcessIot(ctx context.Context, req *iot.IotRequest) (*iot.IotResponse, error) {
	l := req.GetLed()
	resp := iot.IotResponse{}

	err := led.TurnLed(l)
	if err != nil {
		return &resp, status.Error(500, "unable to turn led on/off")
	}

	if l.Active {
		resp.Result = "led turned on"
	} else {
		resp.Result = "led turned off"
	}

	return &resp, nil
}
