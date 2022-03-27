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
 * xmlparser.go -- The XML Config Parser for the sippayphone project
 *
 */

package sippayphone

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ConfigStruct struct {
	XMLName xml.Name `xml:"document"`
	Global  struct {
		Software struct {
			Settings struct {
				SingleInstance     bool   `xml:"singleinstance"`
				OutputDevice       string `xml:"outputdevice"`
				LogFilenameAndPath string `xml:"logfilenameandpath"`
				Logging            string `xml:"logging"`
				Loglevel           string `xml:"loglevel"`
			} `xml:"settings"`
			Sounds struct {
				Sound []struct {
					Event    string `xml:"event,attr"`
					File     string `xml:"file,attr"`
					Volume   string `xml:"volume,attr"`
					Blocking bool   `xml:"blocking,attr"`
					Enabled  bool   `xml:"enabled,attr"`
				} `xml:"sound"`
				Input struct {
					Enabled bool `xml:"enabled,attr"`
					Sound   []struct {
						Event   string `xml:"event,attr"`
						File    string `xml:"file,attr"`
						Enabled bool   `xml:"enabled,attr"`
					} `xml:"sound"`
				} `xml:"input"`
			} `xml:"sounds"`
			Commands struct {
				Command []struct {
					Action  string `xml:"action,attr"`
					Message string `xml:"message,attr"`
					Enabled bool   `xml:"enabled,attr"`
				} `xml:"command"`
			} `xml:"commands"`
			PrintVariables struct {
				PrintAccount bool `xml:"printaccount"`
			} `xml:"printvariables"`
			GPIO struct {
				Name    string `xml:"name,attr"`
				Enabled bool   `xml:"enabled,attr"`
			} `xml:"gpio"`
		} `xml:"software"`
		Hardware struct {
			TargetBoard string `xml:"targetboard,attr"`
			IO          struct {
				Pins struct {
					Pin []struct {
						Direction string `xml:"direction,attr"`
						Device    string `xml:"device,attr"`
						Name      string `xml:"name,attr"`
						PinNo     uint   `xml:"pinno,attr"`
						Type      string `xml:"type,attr"`
						ID        int    `xml:"chipid,attr"`
						Enabled   bool   `xml:"enabled,attr"`
					} `xml:"pin"`
				} `xml:"pins"`
				VolumeButtonStep struct {
					VolUpStep   int `xml:"volupstep"`
					VolDownStep int `xml:"voldownstep"`
				} `xml:"volumebuttonstep"`
			} `xml:"io"`
			HeartBeat struct {
				Enabled     bool   `xml:"enabled,attr"`
				LEDPin      string `xml:"heartbeatledpin"`
				Periodmsecs int    `xml:"periodmsecs"`
				LEDOnmsecs  int    `xml:"ledonmsecs"`
				LEDOffmsecs int    `xml:"ledoffmsecs"`
			} `xml:"heartbeat"`
			LCD struct {
				Enabled               bool   `xml:"enabled,attr"`
				InterfaceType         string `xml:"lcdinterfacetype"`
				I2CAddress            uint8  `xml:"lcdi2caddress"`
				BacklightTimerEnabled bool   `xml:"lcdbacklighttimerenabled"`
				BackLightTimeoutSecs  int    `xml:"lcdbacklighttimeoutsecs"`
				BackLightLEDPin       string `xml:"lcdbacklightpin"`
				RsPin                 int    `xml:"lcdrspin"`
				EPin                  int    `xml:"lcdepin"`
				D4Pin                 int    `xml:"lcdd4pin"`
				D5Pin                 int    `xml:"lcdd5pin"`
				D6Pin                 int    `xml:"lcdd6pin"`
				D7Pin                 int    `xml:"lcdd7pin"`
			} `xml:"lcd"`
			USBKeyboard struct {
				Enabled         bool   `xml:"enabled,attr"`
				USBKeyboardPath string `xml:"usbkeyboarddevpath"`
				NumlockScanID   rune   `xml:"numlockscanid"`
			} `xml:"usbkeyboard"`
			Keyboard struct {
				Command []struct {
					Action      string `xml:"action,attr"`
					ParamName   string `xml:"paramname,attr"`
					Paramvalue  string `xml:"paramvalue,attr"`
					Enabled     bool   `xml:"enabled,attr"`
					Ttykeyboard struct {
						Scanid   rune   `xml:"scanid,attr"`
						Keylabel uint32 `xml:"keylabel,attr"`
						Enabled  bool   `xml:"enabled,attr"`
					} `xml:"ttykeyboard"`
					Usbkeyboard struct {
						Scanid   rune   `xml:"scanid,attr"`
						Keylabel uint32 `xml:"keylabel,attr"`
						Enabled  bool   `xml:"enabled,attr"`
					} `xml:"usbkeyboard"`
				} `xml:"command"`
			} `xml:"keyboard"`
		} `xml:"hardware"`
		Multimedia struct {
			Media struct {
				Source []struct {
					Name     string  `xml:"name,attr"`
					File     string  `xml:"file,attr"`
					Volume   int     `xml:"volume,attr"`
					Duration float32 `xml:"duration,attr"`
					Offset   float32 `xml:"offset,attr"`
					Loop     int     `xml:"loop,attr"`
					Blocking bool    `xml:"blocking"`
					Enabled  bool    `xml:"enabled,attr"`
				} `xml:"source"`
			} `xml:"media"`
		} `xml:"multimedia"`
	} `xml:"global"`
}

type KBStruct struct {
	Enabled    bool
	KeyLabel   uint32
	Command    string
	ParamName  string
	ParamValue string
}

type EventSoundStruct struct {
	Enabled  bool
	FileName string
	Volume   string
	Blocking bool
}

type InputEventSoundFileStruct struct {
	Event   string
	File    string
	Enabled bool
}

// Generic Global Config Variables
var Config ConfigStruct
var ConfigXMLFile string

// Generic Global State Variables
var (
	LCDIsDark bool
)

// Generic Global Timer Variables
var (
	BackLightTime    = time.NewTicker(5 * time.Second)
	BackLightTimePtr = &BackLightTime
	StartTime        = time.Now()
)

var (
	LcdText   = [4]string{"nil", "nil", "nil", "nil"}
	TTYKeyMap = make(map[rune]KBStruct)
	USBKeyMap = make(map[rune]KBStruct)
)

//HD44780 LCD Screen Settings Golbal Variables
var (
	LCDEnabled               bool
	LCDInterfaceType         string
	LCDI2CAddress            uint8
	LCDBackLightTimerEnabled bool
	LCDBackLightTimeout      time.Duration
	LCDRSPin                 int
	LCDEPin                  int
	LCDD4Pin                 int
	LCDD5Pin                 int
	LCDD6Pin                 int
	LCDD7Pin                 int
)

func readxmlconfig(file string, reloadxml bool) error {
	var ReConfig ConfigStruct

	xmlFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	log.Println("info: Successfully Read file " + filepath.Base(file))
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	if !reloadxml {
		err = xml.Unmarshal(byteValue, &Config)
		if err != nil {
			return fmt.Errorf(filepath.Base(file) + " " + err.Error())
		}
	} else {
		err = xml.Unmarshal(byteValue, &ReConfig)
		if err != nil {
			return fmt.Errorf(filepath.Base(file) + " " + err.Error())
		}
	}
	CheckConfigSanity(reloadxml)

	for _, kMainCommands := range Config.Global.Hardware.Keyboard.Command {
		if kMainCommands.Enabled {
			if kMainCommands.Ttykeyboard.Enabled {
				TTYKeyMap[kMainCommands.Ttykeyboard.Scanid] = KBStruct{kMainCommands.Ttykeyboard.Enabled, kMainCommands.Ttykeyboard.Keylabel, kMainCommands.Action, kMainCommands.ParamName, kMainCommands.Paramvalue}
			}
			if kMainCommands.Usbkeyboard.Enabled {
				USBKeyMap[kMainCommands.Usbkeyboard.Scanid] = KBStruct{kMainCommands.Usbkeyboard.Enabled, kMainCommands.Usbkeyboard.Keylabel, kMainCommands.Action, kMainCommands.ParamName, kMainCommands.Paramvalue}
			}

		}
	}

	if strings.ToLower(Config.Global.Software.Settings.Logging) != "screen" && Config.Global.Software.Settings.LogFilenameAndPath == "" {
		Config.Global.Software.Settings.LogFilenameAndPath = "/var/log/"
	}

	LCDEnabled = Config.Global.Hardware.LCD.Enabled
	LCDInterfaceType = Config.Global.Hardware.LCD.InterfaceType
	LCDI2CAddress = Config.Global.Hardware.LCD.I2CAddress
	LCDBackLightTimerEnabled = Config.Global.Hardware.LCD.Enabled
	LCDBackLightTimeout = time.Duration(Config.Global.Hardware.LCD.BackLightTimeoutSecs)
	LCDRSPin = Config.Global.Hardware.LCD.RsPin
	LCDEPin = Config.Global.Hardware.LCD.EPin
	LCDD4Pin = Config.Global.Hardware.LCD.D4Pin
	LCDD5Pin = Config.Global.Hardware.LCD.D5Pin
	LCDD6Pin = Config.Global.Hardware.LCD.D6Pin
	LCDD7Pin = Config.Global.Hardware.LCD.D7Pin

	if Config.Global.Hardware.TargetBoard != "rpi" {
		LCDBackLightTimerEnabled = false
	}

	log.Println("info: Successfully loaded XML configuration file into memory")

	Config.Global.Software.PrintVariables = ReConfig.Global.Software.PrintVariables
	Config.Global.Hardware.Keyboard.Command = ReConfig.Global.Hardware.Keyboard.Command
	Config.Global.Multimedia = ReConfig.Global.Multimedia
	return nil
}

func printxmlconfig() {
	return
}

func CheckConfigSanity(reloadxml bool) {

	Warnings := 0
	Alerts := 0

	log.Println("info: Starting XML Configuration Sanity and Logical Checks")

	if Warnings+Alerts > 0 {
		if Alerts > 0 {
			FatalCleanUp("alert: Fatal Errors Found In sippayphone.xml config file please fix errors, sippayphone stopping now!")
		}

		if Warnings > 0 {
			log.Println("warn: Non-Critical Errors Found In sippayphone.xml config file please fix errors or sippayphone may not behave as expected")
		}
	} else {
		log.Println("info: Finished XML Configuration Sanity and Logical Checks Without Any Alerts/Errors/Warnings")
	}
}
