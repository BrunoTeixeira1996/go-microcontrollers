package programs

import (
	"machine"
)

// Blinking LEDs
func BlinkingLEDs() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		Sleep1Sec()

		led.High()
		Sleep1Sec()
	}
}
