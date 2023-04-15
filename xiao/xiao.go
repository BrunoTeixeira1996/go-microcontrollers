package main

import (
	"machine/usb/hid/keyboard"
	"time"
)

func main() {
	kb := keyboard.Port()

	for {
		kb.Write([]byte("brun0testing\n"))
		time.Sleep(2000 * time.Millisecond)
	}
}
