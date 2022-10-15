package led

import (
	"github.com/raykrishardi/iot-controller-grpc/internal/pkg/iot"
	"github.com/stianeikeland/go-rpio"
)

// TurnLed turns LED on/off
func TurnLed(led *iot.Led) error {
	pin := rpio.Pin(led.Pin)

	if err := rpio.Open(); err != nil {
		return err
	}
	defer rpio.Close()

	pin.Output()

	if led.Active {
		pin.High()
	} else {
		pin.Low()
	}

	return nil
}
