package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
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

func main() {
	servAddr := "localhost:4444"
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

	fmt.Println("ready dial 2482 please!")
	readEvents(conn)
}

func readEvents(conn *net.TCPConn) {
	for {
		reply := make([]byte, 1024)
		_, err := conn.Read(reply)

		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		cleanReply := cleanJSON(string(reply), "{", "}")
		fmt.Println("Clean Reply ", cleanReply)

		var PhoneEvent PhoneEventStruct
		json.Unmarshal([]byte(cleanReply), &PhoneEvent)
		fmt.Printf("Event           = %v\n", PhoneEvent.Event)
		fmt.Printf("Type            = %v\n", PhoneEvent.Type)
		fmt.Printf("Class           = %v\n", PhoneEvent.Class)
		fmt.Printf("Accountaor      = %v\n", PhoneEvent.Accountaor)
		fmt.Printf("Direction       = %v\n", PhoneEvent.Direction)
		fmt.Printf("Peeruri         = %v\n", PhoneEvent.Peeruri)
		fmt.Printf("Peerdisplayname = %v\n", PhoneEvent.ID)
		fmt.Printf("ID              = %v\n", PhoneEvent.Peerdisplayname)
		fmt.Printf("Remoteaudiodir  = %v\n", PhoneEvent.Remoteaudiodir)
		fmt.Printf("Remotevideodir  = %v\n", PhoneEvent.Remotevideodir)
		fmt.Printf("Audiodir        = %v\n", PhoneEvent.Audiodir)
		fmt.Printf("Videodir        = %v\n", PhoneEvent.Videodir)
		fmt.Printf("Param           = %v\n", PhoneEvent.Param)
	}
}

func cleanJSON(str, before, after string) string {
	a := strings.SplitAfterN(str, before, 2)
	b := strings.SplitAfterN(a[len(a)-1], after, 2)
	if 1 == len(b) {
		return b[0]
	}
	return "{" + b[0][0:len(b[0])-len(after)] + "}"
}
