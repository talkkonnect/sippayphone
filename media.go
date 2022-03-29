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
 * media.go -- media player for sippayphone
 *
 */

package sippayphone

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

func aplayLocal(fileNameWithPath string) {
	var player string
	var CmdArguments = []string{}

	if path, err := exec.LookPath("aplay"); err == nil {
		CmdArguments = []string{fileNameWithPath, "-q", "-N"}
		player = path
	} else if path, err := exec.LookPath("paplay"); err == nil {
		CmdArguments = []string{fileNameWithPath}
		player = path
	} else {
		return
	}

	log.Printf("debug: player %v CmdArguments %v", player, CmdArguments)

	cmd := exec.Command(player, CmdArguments...)

	_, err := cmd.CombinedOutput()

	if err != nil {
		return
	}
}

func localMediaPlayer(fileNameWithPath string, playbackvolume int, blocking bool, duration float32, loop int) {

	if loop == 0 || loop > 3 {
		log.Println("warn: Infinite Loop or more than 3 loops not allowed")
		return
	}

	CmdArguments := []string{fileNameWithPath, "-volume", strconv.Itoa(playbackvolume), "-autoexit", "-loop", strconv.Itoa(loop), "-autoexit", "-nodisp"}

	if duration > 0 {
		CmdArguments = []string{fileNameWithPath, "-volume", strconv.Itoa(playbackvolume), "-autoexit", "-t", fmt.Sprintf("%.1f", duration), "-loop", strconv.Itoa(loop), "-autoexit", "-nodisp"}
	}

	cmd := exec.Command("/usr/bin/ffplay", CmdArguments...)

	WaitForFFPlay := make(chan struct{})
	go func() {
		cmd.Run()
		if blocking {
			WaitForFFPlay <- struct{}{} // signal that the routine has completed
		}
	}()
	if blocking {
		<-WaitForFFPlay
	}
}

func PlayTone(toneFreq int, toneDuration float32, destination string, withRXLED bool) {

	if destination == "local" {

		cmdArguments := []string{"-f", "lavfi", "-i", "sine=frequency=" + strconv.Itoa(toneFreq) + ":duration=" + fmt.Sprintf("%f", toneDuration), "-autoexit", "-nodisp"}
		cmd := exec.Command("/usr/bin/ffplay", cmdArguments...)
		var out bytes.Buffer
		cmd.Stdout = &out

		if withRXLED {
			GPIOOutPin("voiceactivity", "on")
		}
		err := cmd.Run()
		if err != nil {
			log.Println("error: ffplay error ", err)
			if withRXLED {
				GPIOOutPin("voiceactivity", "off")
			}
			return
		}
		if withRXLED {
			GPIOOutPin("voiceactivity", "off")
		}

		log.Printf("info: Played Tone at Frequency %v Hz With Duration of %v Seconds\n", toneFreq, toneDuration)
	}
}

func findEventSound(findEventSound string) EventSoundStruct {
	for _, sound := range Config.Global.Software.Sounds.Sound {
		if sound.Enabled && sound.Event == findEventSound {
			return EventSoundStruct{sound.Enabled, sound.File, sound.Volume, sound.Blocking}
		}
	}
	return EventSoundStruct{false, "", "0", false}
}

func findInputEventSoundFile(findInputEventSound string) InputEventSoundFileStruct {
	for _, sound := range Config.Global.Software.Sounds.Input.Sound {
		if sound.Event == findInputEventSound {
			if sound.Enabled {
				return InputEventSoundFileStruct{sound.Event, sound.File, sound.Enabled}
			}
		}
	}
	return InputEventSoundFileStruct{findInputEventSound, "", false}
}

func playIOMedia(inputEvent string) {
	if Config.Global.Software.Sounds.Input.Enabled {
		var inputEventSoundFile InputEventSoundFileStruct = findInputEventSoundFile(inputEvent)
		if inputEventSoundFile.Enabled {
			go aplayLocal(inputEventSoundFile.File)
		}
	}
}
