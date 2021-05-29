package encdec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
)

// convert text password to sha256sum bytes
func newKeyBytes(textPasswd string) []byte {
	hash := sha256.Sum256([]byte(textPasswd))
	return hash[:]
}

//
func EncryptToBase64(msg string, textPasswd string) (string, error) {
	encryptedBytes, err := EncryptToBytes(msg, textPasswd)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(encryptedBytes), nil
}

//
func DecryptFromBase64(encryptedBase64 string, textPasswd string) (string, error) {
	nullResult := ""
	encryptedBytes, err := base64.URLEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return nullResult, err
	}
	result, err := DecryptFromBytes(encryptedBytes, textPasswd)
	return string(result), err
}

// encrypt msg string to []byte
func EncryptToBytes(msg string, textPasswd string) ([]byte, error) {
	nullResult := []byte{}
	keyBytes := newKeyBytes(textPasswd)
	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nullResult, err
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nullResult, err
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nullResult, err
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	cipherText := aesGCM.Seal(nonce, nonce, []byte(msg), nil)
	//fmt.Println("cipherText: ", cipherText, "  ", hex.EncodeToString(cipherText))
	return cipherText, nil
}

// 解密一个[]byte, 解密后的结果也是[]byte
func DecryptFromBytes(encryptedBytes []byte, textPasswd string) ([]byte, error) {
	keyBytes := newKeyBytes(textPasswd)
	nullResult := []byte{}
	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nullResult, err
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()
	//Extract the nonce from the encrypted data
	nonce, cipherText := encryptedBytes[:nonceSize], encryptedBytes[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Println(err)
		return nullResult, err
	}
	return plaintext, nil
}
