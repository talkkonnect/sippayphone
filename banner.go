package sippayphone

import (
	"log"
	"strconv"
)

func sippayphoneBanner(backgroundcolor string) {
	var backgroundreset string = "\u001b[0m"
	log.Println("info: " + backgroundcolor + "┌────────────────────────────────────────────────────────────────┐" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│___(_)_ __  _ __   __ _ _   _ _ __ | |__   ___  _ __   ___      │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│/ __| | '_ \\| '_ \\ / _` | | | | '_ \\| '_ \\ / _ \\| '_ \\ / _ \\    │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│\\__ \\ | |_) | |_) | (_| | |_| | |_) | | | | (_) | | | |  __/    │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│|___/_| .__/| .__/ \\__,_|\\__, | .__/|_| |_|\\___/|_| |_|\\___|    │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│      |_|   |_|          |___/|_|                               │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "├────────────────────────────────────────────────────────────────┤" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│A SIP Phone Using Raspberry Pi Housed in Vintage TOT Pay Phone  │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "├────────────────────────────────────────────────────────────────┤" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│Created By : Suvir Kumar  <suvir@talkkonnect.com>               │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "├────────────────────────────────────────────────────────────────┤" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│Press the <Del> key for Menu or <Ctrl-c> to Quit talkkonnect    │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│Released under MPL 2.0 License                                  │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│Blog at www.talkkonnect.com, source at github.com/talkkonnect   │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "└────────────────────────────────────────────────────────────────┘" + backgroundreset)
	log.Printf("info: Software Ver %v Rel %v \n", sippayphoneVersion, sippayphoneReleased)
	boardVersion := checkSBCVersion()
	if boardVersion != "unknown" {
		log.Printf("info: Hardware Detected As %v\n", boardVersion)
	} else {
		log.Println("info: ")
	}
}

func sippayphoneAcknowledgements(backgroundcolor string) {
	var backgroundreset string = "\u001b[0m"
	log.Println("info: " + backgroundcolor + "┌──────────────────────────────────────────────────────────────────────────────────────────────┐" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│Acknowledgements & Inspriation from the aippayphone team of developers, maintainers & testers │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│sippayphone is based on the works of many people and many open source projects                │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "├──────────────────────────────────────────────────────────────────────────────────────────────┤" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│Thanks to Organizations :-                                                                    │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│The BareSIP Developmentg team, Raspberry Pi Foundation, Developers and Maintainers of Debian  │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│The Creators and Maintainers of Golang and all the libraries available on github.com          │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│Global Coders Co., Ltd. For Sponsoring this project                                           │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "├──────────────────────────────────────────────────────────────────────────────────────────────┤" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│visit us at www.talkkonnect.com and github.com/talkkonnect                                    │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│sippayphone was created by Suvir Kumar <suvir@talkkonnect.com> & Released under MPLV2 License │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "└──────────────────────────────────────────────────────────────────────────────────────────────┘" + backgroundreset)
}

func sippayphoneMenu(backgroundcolor string) {
	var backgroundreset string = "\u001b[0m"
	log.Println("info: " + backgroundcolor + "┌──────────────────────────────────────────────────────────────┐" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│     _ __ ___   __ _(_)_ __    _ __ ___   ___ _ __  _   _     │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│    | '_ ` _ \\ / _` | | '_ \\  | '_ ` _ \\ / _ \\ '_ \\| | | |    │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│    | | | | | | (_| | | | | | | | | | | |  __/ | | | |_| |    │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│    |_| |_| |_|\\__,_|_|_| |_| |_| |_| |_|\\___|_| |_|\\__,_|    │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "├─────────────────────────────┬────────────────────────────────┤" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│ <Del> to Display this Menu  | <Ctrl-C> to Quit sippayphone   │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "├─────────────────────────────┼────────────────────────────────┤" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│ <F1>  Channel Up (+)        │ <F2>  Channel Down (-)         │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│ <F3>  Mute/Unmute Speaker   │ <F4>  Current Volume Level     │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│ <F5>  Digital Volume Up (+) │ <F6>  Digital Volume Down (-)  │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│ <F7>  List Server Channels  │ <F8>  Start Transmitting       │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│ <F9>  Stop Transmitting     │ <F10> List Online Users        │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│ <F11> Playback/Stop Stream  │ <F12> For GPS Position         │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "├─────────────────────────────┼────────────────────────────────┤" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│<Ctrl-B> Reload XML Config   │ <Ctrl-C> Stop Talkkonnect      │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│<Ctrl-D> Debug Stacktrace    │ <Ctrl-E> Send Email            │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "├─────────────────────────────┼────────────────────────────────┤" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│<Ctrl-F> Conn Previous Server│<Ctrl-G> Send Repeater Tone     │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│<Ctrl-H> XML Config Checker  │<Ctrl-I> Traffic Record         │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│<Ctrl-J> Mic Record          │<Ctrl-K> Traffic & Mic Record   │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│<Ctrl-L> Clear Screen        │<Ctrl-M> Radio Channel (+)      │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│<Ctrl-N> Next Server         │<Ctrl-O> Ping Servers           │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│<Ctrl-P> Panic Simulation    │<Ctrl-R> Repeat TX Loop Test    │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│<Ctrl-S> Scan Channels       │<Ctrl-T> Thanks/Acknowledgements│" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│<Ctrl-U> Show Uptime         │<Ctrl-V> Display Version        │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│<Ctrl-X> Dump XML Config     │                                │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "├─────────────────────────────┼────────────────────────────────┤" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│  Visit us at www.talkkonnect.com and github.com/talkkonnect  │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "│  Thanks to Global Coders Co., Ltd. for their sponsorship     │" + backgroundreset)
	log.Println("info: " + backgroundcolor + "└──────────────────────────────────────────────────────────────┘" + backgroundreset)
	log.Println("info: IP Address & Session Information")
	localAddresses()

	macaddress, err := getMacAddr()
	if err != nil {
		log.Println("error: Could Not Get Network Interface MAC Address")
	} else {
		for i, a := range macaddress {
			log.Println("info: Network Interface MAC Address (" + strconv.Itoa(i) + "): " + a)
		}
	}
}
