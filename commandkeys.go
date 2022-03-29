package sippayphone

import (
	"log"
	"strconv"
	"time"

	"github.com/talkkonnect/volume-go"
)

func cmdDisplayMenu() {
	log.Println("debug: Delete Key Pressed Menu and Session Information Requested")

	sippayphoneMenu("\u001b[44;1m") // add blue background to banner reference https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html#background-colors
}

func cmdCurrentVolume() {
	OrigVolume, err := volume.GetVolume(Config.Global.Software.Settings.OutputVolControlDevice)
	if err != nil {
		log.Printf("error: Unable to get current volume: %+v\n", err)
	}

	log.Printf("debug: F4 pressed Volume Level Requested\n")
	log.Println("info: Volume Level is at", OrigVolume, "%")

	if Config.Global.Hardware.TargetBoard == "rpi" {
		if LCDEnabled {
			LcdText = [4]string{"nil", "nil", "nil", "Volume " + strconv.Itoa(OrigVolume)}
			LcdDisplay(LcdText, LCDRSPin, LCDEPin, LCDD4Pin, LCDD5Pin, LCDD6Pin, LCDD7Pin, LCDInterfaceType, LCDI2CAddress)
		}
	}
}

func cmdVolumeUp() {
	log.Printf("debug: F5 pressed Volume UP (+) \n")
	origVolume, err := volume.GetVolume(Config.Global.Software.Settings.OutputVolControlDevice)
	if err != nil {
		log.Printf("warn: unable to get original volume: %+v volume control will not work!\n", err)
		return
	}

	if origVolume < 100 {
		err := volume.IncreaseVolume(Config.Global.Hardware.IO.VolumeButtonStep.VolUpStep, Config.Global.Software.Settings.OutputVolControlDevice)
		if err != nil {
			log.Println("warn: F5 Increase Volume Failed! ", err)
		}
		origVolume, _ := volume.GetVolume(Config.Global.Software.Settings.OutputVolControlDevice)
		log.Println("info: Volume UP (+) Now At ", origVolume, "%")
		if Config.Global.Hardware.TargetBoard == "rpi" {
			if LCDEnabled {
				LcdText = [4]string{"nil", "nil", "nil", "Volume + " + strconv.Itoa(origVolume)}
				LcdDisplay(LcdText, LCDRSPin, LCDEPin, LCDD4Pin, LCDD5Pin, LCDD6Pin, LCDD7Pin, LCDInterfaceType, LCDI2CAddress)
			}
		}
	} else {
		log.Println("debug: F5 Increase Volume")
		log.Println("info: Already at Maximum Possible Volume")
		if Config.Global.Hardware.TargetBoard == "rpi" {
			if LCDEnabled {
				LcdText = [4]string{"nil", "nil", "nil", "Max Vol"}
				LcdDisplay(LcdText, LCDRSPin, LCDEPin, LCDD4Pin, LCDD5Pin, LCDD6Pin, LCDD7Pin, LCDInterfaceType, LCDI2CAddress)
			}
		}
	}
}

func cmdVolumeDown() {
	log.Printf("info: F6 pressed Volume Down (-) \n")
	origVolume, err := volume.GetVolume(Config.Global.Software.Settings.OutputVolControlDevice)
	if err != nil {
		log.Printf("warn: unable to get original volume: %+v volume control will not work!\n", err)
		return
	}

	if origVolume > 0 {
		err := volume.IncreaseVolume(Config.Global.Hardware.IO.VolumeButtonStep.VolDownStep, Config.Global.Software.Settings.OutputVolControlDevice)
		if err != nil {
			log.Println("error: F6 Decrease Volume Failed! ", err)
		}
		origVolume, _ := volume.GetVolume(Config.Global.Software.Settings.OutputVolControlDevice)
		log.Println("info: Volume Down (-) Now At ", origVolume, "%")
		if Config.Global.Hardware.TargetBoard == "rpi" {
			if LCDEnabled {
				LcdText = [4]string{"nil", "nil", "nil", "Volume - " + strconv.Itoa(origVolume)}
				LcdDisplay(LcdText, LCDRSPin, LCDEPin, LCDD4Pin, LCDD5Pin, LCDD6Pin, LCDD7Pin, LCDInterfaceType, LCDI2CAddress)
			}
		} else {
			log.Println("debug: F6 Increase Volume Already")
			log.Println("info: Already at Minimum Possible Volume")
			if Config.Global.Hardware.TargetBoard == "rpi" {
				if LCDEnabled {
					LcdText = [4]string{"nil", "nil", "nil", "Min Vol"}
					LcdDisplay(LcdText, LCDRSPin, LCDEPin, LCDD4Pin, LCDD5Pin, LCDD6Pin, LCDD7Pin, LCDInterfaceType, LCDI2CAddress)
				}
			}
		}
	}
}

func cmdQuitsippayphone() {
	log.Printf("debug: Ctrl-C Terminate Program Requested \n")
	duration := time.Since(StartTime)
	log.Printf("info: sippayphone Now Running For %v \n", secondsToHuman(int(duration.Seconds())))
	CleanUp()
}

func cmdClearScreen() {
	reset()
	log.Printf("debug: Ctrl-L Pressed Cleared Screen \n")
	if Config.Global.Hardware.TargetBoard == "rpi" {
		if LCDEnabled {
			LcdText = [4]string{"nil", "nil", "nil", "nil"}
			LcdDisplay(LcdText, LCDRSPin, LCDEPin, LCDD4Pin, LCDD5Pin, LCDD6Pin, LCDD7Pin, LCDInterfaceType, LCDI2CAddress)
		}
	}
}

func cmdThanks() {
	log.Printf("debug: Ctrl-T Pressed \n")
	log.Println("info: Thanks and Acknowledgements Screen Request ")
	sippayphoneAcknowledgements("\u001b[44;1m") // add blue background to banner reference https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html#background-colors
}

func cmdShowUptime() {
	log.Printf("debug: Ctrl-U Pressed \n")
	log.Println("info: sippayphone Uptime Request ")
	duration := time.Since(StartTime)
	log.Printf("info: sippayphone Now Running For %v \n", secondsToHuman(int(duration.Seconds())))
}

func cmdDisplayVersion() {
	log.Printf("debug: Ctrl-V Pressed \n")
	log.Println("info: sippayphone Version Request ")
	log.Printf("info: Ver %v Rel %v \n", sippayphoneVersion, sippayphoneReleased)
}

func cmdDumpXMLConfig() {
	log.Printf("debug: Ctrl-X Pressed \n")
	log.Println("info: Print XML Config " + ConfigXMLFile)
	printxmlconfig()
}

func cmdLiveReload() {
	log.Printf("debug: Ctrl-B Pressed \n")
	log.Println("info: XML Config Live Reload")
	err := readxmlconfig(ConfigXMLFile, true)
	if err != nil {
		message := err.Error()
		FatalCleanUp(message)
	}
}

func cmdSanityCheck() {
	log.Printf("debug: Ctrl-H Pressed \n")
	log.Println("info: XML Sanity Checker")
	CheckConfigSanity(false)
}
