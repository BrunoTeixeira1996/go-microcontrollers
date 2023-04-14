package main

import (
	"machine/usb/hid/keyboard"
	"time"
)

func xiao() {
	kb := keyboard.Port()

	for {
		kb.Write([]byte("brun0testing\n"))
		time.Sleep(2000 * time.Millisecond)
	}
}

func main() {
	xiao()
}
