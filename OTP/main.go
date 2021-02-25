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

	interactiveFlag := flag.Bool("i", false, "toggle interactive mode")
	modeFlag := flag.String("mode", "e", "(e)ncipher/(d)ecipher (required if not in interactive mode)")
	keyFlag := flag.String("key", "", "Enter an OTP key (required if not in interactive mode)")
	msgFlag := flag.String("message", "", "Enter a message to encipher/decipher (required if not in interactive mode)")
	genKeyFlag := flag.Bool("genkey", false, "Generate a key (must supply a message and mode)")
	flag.Parse()

	if *interactiveFlag {
		interactiveInput()
	} else if *genKeyFlag {
		var ret otp.Message
		scanner := bufio.NewScanner(os.Stdin)
		log.Println("Enter message to encipher:")
		scanner.Scan()
		ret.Message = scanner.Text()
		ret.GenerateKey()
		log.Println("Your message: ", ret.Result)
		log.Println("Your generated key is: ", ret.Key)

	} else {
		var ret otp.Message
		if len(*keyFlag) == len(*msgFlag) && len(*keyFlag) != 0 {
			ret.Key = *keyFlag
			ret.Message = *msgFlag
			if *modeFlag == "e" {
				log.Println(*modeFlag, *keyFlag, *msgFlag)
				ret.Encrypt()
			} else if *modeFlag == "d" {
				log.Println(*modeFlag, *keyFlag, *msgFlag)
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
	}
}
