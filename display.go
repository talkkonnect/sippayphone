package sippayphone

import "github.com/talkkonnect/go-hd44780"

func LcdDisplay(lcdtextshow [4]string, PRSPin int, PEPin int, PD4Pin int, PD5Pin int, PD6Pin int, PD7Pin int, LCDInterfaceType string, LCDI2CAddress byte) {
	go hd44780.LcdDisplay(LcdText, LCDRSPin, LCDEPin, LCDD4Pin, LCDD5Pin, LCDD6Pin, LCDD7Pin, LCDInterfaceType, LCDI2CAddress)
}
