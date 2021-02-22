package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AlphaLU/GoAlgos/OTP/otp"
)

var p = fmt.Println

func main() {
	// encMessage := "J2VUOSMMRG2LL"
	// orgMessage := "GRILLEDCHEESE"
	// key := "DLNJDOJKKCY3H"

	modeFlag := flag.String("mode", "e", "(e)ncipher/(d)ecipher")
	keyFlag := flag.String("key", "", "Enter an OTP key")
	msgFlag := flag.String("message", "", "Enter a message to encipher/decipher")

	flag.Parse()
	log.Println(*modeFlag, *keyFlag, *msgFlag)
	var ret otp.Message
	ret.Key = *keyFlag
	ret.Message = *msgFlag
	if *modeFlag == "e" {
		ret.Encrypt()
	} else {
		ret.Decrypt()
	}

	p(ret.Result)
	// var msg otp.Message
	// msg.Key = key
	// msg.Message = encMessage
	// msg.Decrypt()
	// p(msg.Result)

	// var msg2 otp.Message
	// msg2.Key = key
	// msg2.Message = orgMessage
	// msg2.Encrypt()
	// p(msg2.Result)
}
