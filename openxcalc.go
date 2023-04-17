package main

import (
	"machine/usb/hid/keyboard"
	"time"
)


const (

	GUI = 0x08 | 0xE000
	ENTER = 40 | 0xF000
)

// https://github.com/tinygo-org/tinygo/blob/release/src/machine/usb/hid/keyboard/keycode.go

/*
WINDOWS KEY + ENTER (opens dmenu)
WRITE xcalc
ENTER
*/
func main() {
	kb := keyboard.Port()

	count := 1
	for {
		if count == 2 {
			break
		}
		count +=1

		time.Sleep(1000 * time.Millisecond)
		
		kb.Down(GUI)
		kb.Down(ENTER)
		kb.Release()

		time.Sleep(1000 * time.Millisecond)
		
		kb.Write([]byte("xcalc"))

		time.Sleep(1000 * time.Millisecond)

		kb.Press(ENTER)
	}
}