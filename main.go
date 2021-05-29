package main

import (
	"github.com/leo2n/selfTools/encdec"
	"log"
)

func main() {
	log.Println("Hello Log")
	log.Println(encdec.Encrypt("Hello World"))
}
