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
 * main.go -- The Main Function of sippayphone project
 *
 */

package sippayphone

import (
	"encoding/json"
	"log"
	"net"
	"os"

	"github.com/comail/colog"
	term "github.com/talkkonnect/termbox-go"
)

type PhoneEventStruct struct {
	Event           bool   `json:"event"`
	Type            string `json:"type"`
	Class           string `json:"class"`
	Accountaor      string `json:"accountaor"`
	Direction       string `json:"direction"`
	Peeruri         string `json:"peeruri"`
	Peerdisplayname string `json:"peerdisplayname"`
	ID              string `json:"id"`
	Remoteaudiodir  string `json:"remoteaudiodir"`
	Remotevideodir  string `json:"remotevideodir"`
	Audiodir        string `json:"audiodir"`
	Videodir        string `json:"videodir"`
	Param           string `json:"param"`
}

var (
	servPort   string = "4444"
	servAddr   string = "localhost" + ":" + servPort
	cleanReply string
)

func Init(file string) {
	err := term.Init()
	if err != nil {
		FatalCleanUp("Cannot Initialize Terminal Error: " + err.Error())
	}
	defer term.Close()

	colog.Register()
	colog.SetOutput(os.Stdout)

	ConfigXMLFile = file
	err = readxmlconfig(ConfigXMLFile, false)
	if err != nil {
		message := err.Error()
		FatalCleanUp(message)
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	log.Println("ready dial 2482 to call please !")
	readEvents(conn)
}

func readEvents(conn *net.TCPConn) {
	for {
		reply := make([]byte, 1024)
		_, err := conn.Read(reply)

		if err != nil {
			println("read from server failed:", err.Error())
			os.Exit(1)
		}

		cleanReply = cleanJSON(string(reply), "{", "}")
		log.Println("Clean Reply ", cleanReply)

		var PhoneEvent PhoneEventStruct
		json.Unmarshal([]byte(cleanReply), &PhoneEvent)
		log.Printf("Event           = %v\n", PhoneEvent.Event)
		log.Printf("Type            = %v\n", PhoneEvent.Type)
		log.Printf("Class           = %v\n", PhoneEvent.Class)
		log.Printf("Accountaor      = %v\n", PhoneEvent.Accountaor)
		log.Printf("Direction       = %v\n", PhoneEvent.Direction)
		log.Printf("Peeruri         = %v\n", PhoneEvent.Peeruri)
		log.Printf("Peerdisplayname = %v\n", PhoneEvent.ID)
		log.Printf("ID              = %v\n", PhoneEvent.Peerdisplayname)
		log.Printf("Remoteaudiodir  = %v\n", PhoneEvent.Remoteaudiodir)
		log.Printf("Remotevideodir  = %v\n", PhoneEvent.Remotevideodir)
		log.Printf("Audiodir        = %v\n", PhoneEvent.Audiodir)
		log.Printf("Videodir        = %v\n", PhoneEvent.Videodir)
		log.Printf("Param           = %v\n", PhoneEvent.Param)
	}
}
