package sippayphone

import (
	"log"
	"strings"

	evdev "github.com/gvalkov/golang-evdev"
)

func USBKeyboard() {

	device, err := evdev.Open(Config.Global.Hardware.USBKeyboard.USBKeyboardPath)
	if err != nil {
		log.Printf("error: Unable to open USB Keyboard input device: %s\nError: %v It will now Be Disabled\n", Config.Global.Hardware.USBKeyboard.USBKeyboardPath, err)
		return
	}

	var keyPrevStateDown bool

	for {
		events, err := device.Read()
		if err != nil {
			log.Printf("error: Unable to Read Event From USB Keyboard error %v\n", err)
			return
		}

		for _, ev := range events {
			switch ev.Type {
			case evdev.EV_KEY:
				ke := evdev.NewKeyEvent(&ev)

				if ke.State == evdev.KeyDown {
					keyPrevStateDown = true
				}

				// Functions that we allow Repeating Keys Defined Here
				if ke.State == evdev.KeyHold {
					keyPrevStateDown = false
					if _, ok := USBKeyMap[rune(ke.Scancode)]; ok {
						switch strings.ToLower(USBKeyMap[rune(ke.Scancode)].Command) {
						case "channelup":
							playIOMedia("usbchannelup")
						case "channeldown":
							playIOMedia("usbchanneldown")
						case "volumeup":
							playIOMedia("usbvolup")
						case "volumedown":
							playIOMedia("usbvoldown")
						}
					} else {
						if ke.Scancode != uint16(Config.Global.Hardware.USBKeyboard.NumlockScanID) {
							log.Println("error: Key Not Mapped ASC ", ke.Scancode)
						}
					}
					continue
				}

				//Key Up & Down One Shot
				if keyPrevStateDown && ke.State == evdev.KeyUp {
					keyPrevStateDown = false
					if _, ok := USBKeyMap[rune(ke.Scancode)]; ok {
						switch strings.ToLower(USBKeyMap[rune(ke.Scancode)].Command) {
						case "channelup":
							playIOMedia("usbchannelup")
						case "channeldown":
							playIOMedia("usbchanneldown")
						default:
							log.Println("error: Command Not Defined ", strings.ToLower(USBKeyMap[rune(ke.Scancode)].Command))
						}
					} else {
						if ke.Scancode != uint16(Config.Global.Hardware.USBKeyboard.NumlockScanID) {
							log.Println("error: Key Not Mapped ASC ", ke.Scancode)
						}
					}
				}
			}
		}
	}
}
