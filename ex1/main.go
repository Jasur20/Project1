package main

import (
	//"brt_adapter/db"
	//"brt_adapter/models"
	// "brt_adapter/routes"
	"brt_adapter/routes"
	"brt_adapter/settings"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	//"github.com/gin-gonic/gin"
)

func main() {

	settings.AppSettings=settings.ReadSettings()
	routes.Init()
	// routes.Init()
	// rand.Seed(time.Now().UnixNano())
	// chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
    // "abcdefghijklmnopqrstuvwxyzåäö" +
    // "0123456789")
	// length := 8
	// var b strings.Builder
	// for i := 0; i < length; i++ {
    // b.WriteRune(chars[rand.Intn(len(chars))])
	// }
	// str := b.String() // Например "ExcbsVQs"
	// fmt.Println(str)
	
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
