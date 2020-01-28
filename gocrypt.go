package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

//export Encrypt
// Encrypt obfuscates the given string (payload) using the provided passphrase
func Encrypt(payload, passPhrase *C.char) *C.char {
	return C.CString(goEncrypt(
		C.GoString(payload),
		C.GoString(passPhrase),
	))
}

//export Decrypt
// Decrypt deobfuscates the real string from a given encrypted string and passphrase.
func Decrypt(payload, passPhrase *C.char) *C.char {
	return C.CString(goDecrypt(
		C.GoString(payload),
		C.GoString(passPhrase),
	))
}

func goEncrypt(payload, passPhrase string) string {
	nonce := []byte(hashString(passPhrase))[:12]
	return encryptWithNonce(payload, passPhrase, nonce)
}

func goDecrypt(payload, passPhrase string) string {
	gKey := keyGen(passPhrase)
	key := base64.StdEncoding.EncodeToString(gKey)
	kb, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	dVal, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return string(decryptAES128GCM(kb, 12, dVal, nil))
}

func encryptWithNonce(data, passPhrase string, nonce []byte) string {
	gKey := keyGen(passPhrase)
	key := base64.StdEncoding.EncodeToString(gKey)
	kb, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	enc := encryptAES128GCMWithNonce(kb, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(enc)
}

func hashString(payload string) string {
	h := sha256.New()
	_, err := h.Write([]byte(payload))
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// if the key is less than 16 bytes, buffer with a known string
const keyString = `rGjoUXnCZiqQ2X00`

func keyGen(passPhrase string) []byte {
	var gk string
	if len(passPhrase) < 16 {
		i := 16 - len(passPhrase)
		gk = passPhrase + keyString[:i]
	} else {
		gk = passPhrase
	}
	return []byte(gk)
}

func encryptAES128GCMWithNonce(key []byte, nonce []byte, encPayload []byte, metadata []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return gcm.Seal(nonce, nonce, encPayload, metadata)
}

func decryptAES128GCM(key []byte, nonceSize int, encPayload []byte, metadata []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	nonce := make([]byte, nonceSize)
	copy(nonce, encPayload)
	out, err := gcm.Open(nil, nonce, encPayload[nonceSize:], metadata)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return out
}

func main() {}
