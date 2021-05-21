package otp

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Message struct {
	Key     string
	Message string
	Result  string
}

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", " "}

//Encipher
func (m *Message) Encrypt() {
	if !CheckValid(m) {
		log.Println("Error, key or message are missing")
	}

	fmt.Println(alphabet, m.Message, m.Key)
	var ab []string
	for _, k := range m.Key {
		for i, c := range alphabet {
			if string(k) == c {
				ab = append(ab, alphabet[i:]...)
				ab = append(ab, alphabet[:i]...)
				fmt.Println(ab, string(k))

				m.Result += getChar(ab, m.Result, m.Message)

				ab = nil
				break
			}
		}
	}
}

//Helper method for Encrypt()
func getChar(arr []string, res, msg string) string {
	for j, v := range alphabet {
		if string(msg[len(res)]) == v {
			return arr[j]

		}
	}
	return ""
}

//Decipher a message
func (m *Message) Decrypt() {
	if !CheckValid(m) {
		log.Println("Error, key or message are missing")
	}
	for k, v := range m.Key {
		for i, l := range alphabet {
			if l == string(v) {
				var ab []string
				ab = append(ab, alphabet[i:]...)
				ab = append(ab, alphabet[:i]...)
				fmt.Println(ab, k)

				for a, r := range ab {
					if string(m.Message[k]) == r {
						m.Result += alphabet[a]
					}
				}

				ab = nil
			}
		}
	}
}

func (m *Message) GenerateKey() {
	keyLength := len(m.Message)
	var seed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, keyLength)
	for i := range b {
		b[i] = charset[seed.Intn(len(charset))]
	}
	m.Key = string(b)
	m.Encrypt()

}

func CheckValid(m *Message) bool {
	if len(m.Key) < 1 {
		return false
	} else if len(m.Message) < 1 {
		return false
	}
	return true
}
