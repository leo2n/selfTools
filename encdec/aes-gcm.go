package encdec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

// Generate 32 byte sha256sum
func newSHA256(origin string) []byte {
	hash := sha256.Sum256([]byte(origin))
	return hash[:]
}

// input: msg string, output: base64([]byte)
func Encrypt(msg string) string {
	log.Println("keyString is:", keyString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(keyString)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	log.Println("nonce is: ", nonce)

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	cipherText := aesGCM.Seal(nonce, nonce, []byte(msg), nil)
	//fmt.Println("cipherText: ", cipherText, "  ", hex.EncodeToString(cipherText))
	return base64.URLEncoding.EncodeToString(cipherText)
}

// 解密: base64([]byte) -> string
func Decrypt(encryptedString string) string {
	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(keyString)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	encrypted, err := base64.URLEncoding.DecodeString(encryptedString)
	if err != nil {
		log.Fatalln("base64 decoding failed")
	}
	//Extract the nonce from the encrypted data
	nonce, cipherText := encrypted[:nonceSize], encrypted[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
