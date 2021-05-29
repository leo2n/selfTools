package main

import (
	"log"

	"github.com/leo2n/selfTools/encdec"
	"github.com/leo2n/selfTools/randLen"
)

func main() {
	log.Println(encdec.Encrypt("Hello World"))
	log.Println(randLen.RandLenStringGenerator(8))
}
