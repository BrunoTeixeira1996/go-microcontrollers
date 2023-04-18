package programs

import (
	"machine/usb/hid/keyboard"
)

// https://github.com/tinygo-org/tinygo/blob/release/src/machine/usb/hid/keyboard/keycode.go

/*
WINDOWS KEY + ENTER (opens dmenu)
WRITE xcalc
ENTER
*/
func Openxcalc() {
	
	kb := keyboard.Port()

	runs := 1
	for {
		if runs == 2 {
			break
		}
		runs +=1

		Sleep1Sec()

		kb.Down(keyboard.KeyModifierGUI)
		kb.Down(keyboard.KeyEnter)
		kb.Release()

		Sleep1Sec()

		kb.Write([]byte("xcalc"))

		Sleep1Sec()

		kb.Press(keyboard.KeyEnter)
	}
}