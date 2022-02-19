package main

import (
	"machine"
	"time"
)

func main() {
	//led := machine.LED
	//led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	//
	//for {
	//	led.Low()
	//	time.Sleep(time.Millisecond * 500)
	//
	//	led.High()
	//	time.Sleep(time.Millisecond * 500)
	//}

	button := machine.D8
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	led := machine.D13
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.Set(!button.Get())
		time.Sleep(time.Millisecond * 100)
	}
}
