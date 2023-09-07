package main

import (
	"errors"
	"flag"

	"github.com/jesprna/nfc-reader/char"
	"github.com/jesprna/nfc-reader/service"
)

func main() {
	var appFlags service.Flags
	var endChar, inChar string
	var ok bool
	//Read application flags
	flag.StringVar(&endChar, "end-char", "none", "Character at the end of UID. Options: "+char.CharFlagOptions())
	flag.StringVar(&inChar, "in-char", "none", "Сharacter between bytes of UID. Options: "+char.CharFlagOptions())
	flag.BoolVar(&appFlags.CapsLock, "caps-lock", false, "UID with Caps Lock")
	flag.BoolVar(&appFlags.Reverse, "reverse", false, "UID reverse order")
	flag.BoolVar(&appFlags.Decimal, "decimal", false, "UID in decimal format")
	flag.IntVar(&appFlags.Device, "device", 0, "Device number to use")
	flag.Parse()

	//Check flags
	appFlags.EndChar, ok = char.StringToCharFlag(endChar)
	if !ok {
		service.ErrorExit(errors.New("Unknown end character flag. Run with '-h' flag to check options"))
		return
	}
	appFlags.InChar, ok = char.StringToCharFlag(inChar)
	if !ok {
		service.ErrorExit(errors.New("Unknown in character flag. Run with '-h' flag to check options"))
		return
	}

	service := service.NewService(appFlags)
	service.Start()

}
