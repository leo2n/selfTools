package main

import (
	"log"

	"github.com/leo2n/selfTools/encdec"
)

func main() {
	textPasswd := ""
	s, _ := encdec.EncryptToBytes("Hello World", textPasswd)

	log.Println(s, "\n", string(s))
	ss, _ := encdec.DecryptFromBytes(s, textPasswd)
	log.Println(string(ss))
}
