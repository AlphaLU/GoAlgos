package main

import (
	"otpgo/otp"
	"testing"
)

var decMessage = "GRILLEDCHEESE"
var key = "DLNJDOJKKCY3H"
var encMessage = "J2VUOSMMRG2LL"

func TestOTP(t *testing.T) {
	var m otp.Message
	m.Key = key
	m.Message = decMessage
	m.Encrypt()
	if m.Result == encMessage {
		t.Logf("Success!! Encryption failed for: %s With key: %s, Got %s, Expected: %s", decMessage, key, m.Result, encMessage)
	}

	m.Key = key
	m.Message = encMessage
	m.Result = ""
	m.Decrypt()
	if m.Result == decMessage {
		t.Logf("Success!!! Decryption failed for: %s With key: %s, Got %s, Expected: %s", encMessage, key, m.Result, decMessage)
	}
}
