package main

import (
	"brt_adapter/routes"
	"brt_adapter/settings"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)


func main() {

	settings.AppSettings=settings.ReadSettings()
	routes.Init()
}

func GetSha256(text string, secret []byte) string {

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, secret)

	// Write Data to it
	h.Write([]byte(text))

	// Get result and encode as hexadecimal string
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}
