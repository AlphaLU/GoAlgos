package main

import (
	"fmt"

	"github.com/AlphaLU/GoAlgos/OTP/otp"
)

var p = fmt.Println

func main() {
	encMessage := "J2VUOSMMRG2LL"
	orgMessage := "GRILLEDCHEESE"
	key := "DLNJDOJKKCY3H"

	var msg otp.Message
	msg.Key = key
	msg.Message = encMessage
	msg.Decrypt()
	p(msg.Result)

	var msg2 otp.Message
	msg2.Key = key
	msg2.Message = orgMessage
	msg2.Encrypt()
	p(msg2.Result)
}
