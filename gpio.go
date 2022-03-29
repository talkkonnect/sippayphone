/*
 * sippayphone headless sip phone with lcd screen and keyboard
 * Copyright (C) 2018-2019, Suvir Kumar <suvir@talkkonnect.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 * Software distributed under the License is distributed on an "AS IS" basis,
 * WITHOUT WARRANTY OF ANY KIND, either express or implied. See the License
 * for the specific language governing rights and limitations under the
 * License.
 *
 * sippayphone is using tcp API json calls to baresip softphone
 * baresip can be found at https://github.com/baresip/baresip
 *
 * The Initial Developer of the Original Code is
 * Suvir Kumar <suvir@talkkonnect.com>
 * Portions created by the Initial Developer are Copyright (C) Suvir Kumar. All Rights Reserved.
 *
 * Contributor(s):
 *
 * Suvir Kumar <suvir@talkkonnect.com>
 *
 * My Blog is at www.talkkonnect.com
 * The source code is hosted at github.com/talkkonnect/sippayphone
 *
 * gpio.go -- gpio hanler for sipppayphone project for raspberry pi
 *
 */

package sippayphone

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
	"github.com/talkkonnect/gpio"
)

//Variables for Input Buttons/Switches
var (
	TxButtonUsed  bool
	TxButton      gpio.Pin
	TxButtonPin   uint
	TxButtonState uint

	VolUpButtonUsed  bool
	VolUpButton      gpio.Pin
	VolUpButtonPin   uint
	VolUpButtonState uint

	VolDownButtonUsed  bool
	VolDownButton      gpio.Pin
	VolDownButtonPin   uint
	VolDownButtonState uint
	GPIOEnabled        bool
)

func initGPIO() {

	if Config.Global.Hardware.TargetBoard != "rpi" {
		return
	}

	if err := rpio.Open(); err != nil {
		log.Println("error: GPIO Error, ", err)
		GPIOEnabled = false
		return
	}
	GPIOEnabled = true

	//handle inputs on RPI GPIO
	for _, io := range Config.Global.Hardware.IO.Pins.Pin {
		if io.Enabled && io.Direction == "input" && io.Type == "gpio" {
			if io.Name == "txptt" && io.PinNo > 0 {
				log.Printf("debug: GPIO Setup Input Device %v Name %v PinNo %v", io.Device, io.Name, io.PinNo)
				TxButtonPinPullUp := rpio.Pin(io.PinNo)
				TxButtonPinPullUp.PullUp()
				TxButtonUsed = true
				TxButtonPin = io.PinNo
			}
		}

		if VolUpButtonUsed || VolDownButtonUsed {
			rpio.Close()
		}

		if VolUpButtonUsed {
			VolUpButton = gpio.NewInput(VolUpButtonPin)
			go func() {
				for {
					currentState, err := VolUpButton.Read()
					time.Sleep(150 * time.Millisecond)

					if currentState != VolUpButtonState && err == nil {
						VolUpButtonState = currentState

						if VolUpButtonState == 1 {
							log.Println("debug: Vol UP Button is released")
						} else {
							log.Println("debug: Vol UP Button is pressed")
							playIOMedia("iovolup")
						}
						time.Sleep(1 * time.Second)
					}
				}
			}()
		}

		if VolDownButtonUsed {
			VolDownButton = gpio.NewInput(VolDownButtonPin)
			go func() {
				for {
					currentState, err := VolDownButton.Read()
					time.Sleep(150 * time.Millisecond)
					if currentState != VolDownButtonState && err == nil {
						VolDownButtonState = currentState
						if VolDownButtonState == 1 {
							log.Println("debug: Vol Down Button is released")
						} else {
							log.Println("debug: Vol Down Button is pressed")
							playIOMedia("iovoldown")
						}
					}
					time.Sleep(1 * time.Second)
				}
			}()
		}
	}
}

func GPIOOutPin(name string, command string) {
	if Config.Global.Hardware.TargetBoard != "rpi" {
		return
	}

	for _, io := range Config.Global.Hardware.IO.Pins.Pin {

		if io.Enabled && io.Direction == "output" && io.Name == name {
			if command == "on" {
				switch io.Type {
				case "gpio":
					log.Printf("debug: Turning On %v at pin %v Output GPIO\n", io.Name, io.PinNo)
					gpio.NewOutput(io.PinNo, true)
				default:
					log.Println("error: GPIO Types Currently Supported are gpio or mcp23017 only!")
				}
				break
			}

			if command == "off" {
				switch io.Type {
				case "gpio":
					log.Printf("debug: Turning Off %v at pin %v Output GPIO\n", io.Name, io.PinNo)
					gpio.NewOutput(io.PinNo, false)
				default:
					log.Println("error: GPIO Types Currently Supported are gpio or mcp23017 only!")
				}
				break
			}

			if command == "pulse" {
				switch io.Type {
				case "gpio":
					log.Printf("debug: Pulsing %v at pin %v Output GPIO\n", io.Name, io.PinNo)
					gpio.NewOutput(io.PinNo, false)
					time.Sleep(1 * time.Millisecond)
					gpio.NewOutput(io.PinNo, true)
					time.Sleep(1 * time.Millisecond)
					gpio.NewOutput(io.PinNo, false)
					time.Sleep(1 * time.Millisecond)
				default:
					log.Println("error: GPIO Types Currently Supported are gpio or mcp23017 only!")
				}
				break
			}
		}
	}
}

func GPIOOutAll(name string, command string) {
	if Config.Global.Hardware.TargetBoard != "rpi" {
		return
	}

	for _, io := range Config.Global.Hardware.IO.Pins.Pin {
		if io.Enabled && io.Direction == "output" && io.Device == "led/relay" {
			switch io.Type {
			case "gpio":
				if command == "on" {
					log.Printf("debug: Turning On %v Output GPIO\n", io.Name)
					gpio.NewOutput(io.PinNo, true)
				}
				if command == "off" {
					log.Printf("debug: Turning Off %v Output GPIO\n", io.Name)
					gpio.NewOutput(io.PinNo, false)
				}
			default:
				log.Println("error: GPIO Types Currently Supported are gpio on board only!")
			}
		}
	}
}
