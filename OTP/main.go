package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AlphaLU/GoAlgos/OTP/otp"
)

var p = fmt.Println

func interactiveInput() {
	var mode string
	var key string
	var msg string

	var ret otp.Message

	scanner := bufio.NewScanner(os.Stdin)

	p("Choose a mode (e)ncipher or (d)ecipher:")
	scanner.Scan()
	mode = scanner.Text()
	if mode == "e" {
		p("Enter the key:")
		scanner.Scan()
		key = scanner.Text()
		if len(key) > 0 {
			p("Enter the message to encrypt:")
			scanner.Scan()
			msg = scanner.Text()
			if len(msg) == len(key) {
				ret.Key = key
				ret.Message = msg
				ret.Encrypt()
				log.Println("Result: ", ret.Result)
			} else {
				p("Key and Message length need to match")
				interactiveInput()
			}
		} else {
			p("Key can't be empty")
			interactiveInput()
		}
	} else if mode == "d" {
		p("Enter the key:")
		scanner.Scan()
		key = scanner.Text()
		if len(key) > 0 {
			p("Enter the message to decrypt:")
			scanner.Scan()
			msg = scanner.Text()
			if len(msg) == len(key) {
				ret.Key = key
				ret.Message = msg
				ret.Decrypt()
				log.Println("Result: ", ret.Result)
			} else {
				p("Key and Message length need to match")
				interactiveInput()
			}
		} else {
			p("Key can't be empty")
			interactiveInput()
		}
	} else {
		p("Bad input...")
		return
	}
}

func main() {
	// encMessage := "J2VUOSMMRG2LL"
	// orgMessage := "GRILLEDCHEESE"
	// key := "DLNJDOJKKCY3H"

	interactiveFlag := flag.Bool("interactive", false, "toggle interactive mode")
	modeFlag := flag.String("mode", "e", "(e)ncipher/(d)ecipher")
	keyFlag := flag.String("key", "", "Enter an OTP key")
	msgFlag := flag.String("message", "", "Enter a message to encipher/decipher")
	log.Println(*interactiveFlag)

	if *interactiveFlag {
		interactiveInput()
	} else {
		flag.Parse()
		log.Println(*modeFlag, *keyFlag, *msgFlag)
		var ret otp.Message
		if len(*keyFlag) == len(*msgFlag) && len(*keyFlag) != 0 {
			ret.Key = *keyFlag
			ret.Message = *msgFlag
			if *modeFlag == "e" {
				ret.Encrypt()
			} else if *modeFlag == "d" {
				ret.Decrypt()
			} else {
				flag.PrintDefaults()
				return
			}
		} else {
			log.Println("The message and the key must be the same size and can't be empty")
			flag.PrintDefaults()
			return
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
}
