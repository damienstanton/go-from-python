package main

import (
	"encoding/base64"
	"testing"
)

const passPhrase = "thisisnotagoodpassphrase"
const payload = "plaintext"

func TestHash(t *testing.T) {
	expected := "0VlIELTI4/YhVdatDKPyD8RikMSR4GBkBixjqccQmCk="
	if hashString(passPhrase) != expected {
		t.Fatalf("expected %s, got %s", expected, hashString(passPhrase))
	}
}

func TestKeyGen(t *testing.T) {
	b := keyGen(passPhrase)
	if string(b) != passPhrase {
		t.Fatal("passphrase not used correctly in keygen")
	}
	short := keyGen("abc")
	appended := "abcrGjoUXnCZiqQ2"
	if string(short) != appended {
		t.Fatal("passphrase not used correctly in keygen")
	}
}

func TestGCMOps(t *testing.T) {
	key := keyGen(passPhrase)
	kb, err := base64.StdEncoding.DecodeString(hashString(string(key)))
	if err != nil {
		t.Fatalf("b64 decode error: %v", err)
	}

	payloadBytes := []byte(payload)
	nonce := []byte(hashString(passPhrase))[:12]

	enc := encryptAES128GCMWithNonce(
		kb,
		nonce,
		key,
		payloadBytes,
	)

	dec := decryptAES128GCM(
		kb,
		12,
		enc,
		payloadBytes,
	)
	if string(dec) != passPhrase {
		t.Fatal("enc/dec error in GCM")
	}
}

func TestPrivateEnc(t *testing.T) {
	nonce := []byte(hashString(passPhrase))[:12]
	enc := encryptWithNonce(payload, passPhrase, nonce)
	expected := "MFZsSUVMVEk0L1loLSZukGqznR0xu2m4DfSfa+b1HoPVYKC2KQ=="
	if enc != expected {
		t.Fatal("encryption error")
	}
}

func TestEncrypt(t *testing.T) {
	enc := goEncrypt(payload, passPhrase)
	expected := "MFZsSUVMVEk0L1loLSZukGqznR0xu2m4DfSfa+b1HoPVYKC2KQ=="
	if enc != expected {
		t.Fatal("encryption error")
	}
}

func TestDecrypt(t *testing.T) {
	enc := "MFZsSUVMVEk0L1loLSZukGqznR0xu2m4DfSfa+b1HoPVYKC2KQ=="
	dec := goDecrypt(enc, passPhrase)
	if dec != payload {
		t.Fatal("decryption error")
	}
}
