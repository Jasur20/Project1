package main

import (
	// "brt_adapter/routes"
	// "brt_adapter/settings"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// settings.AppSettings=settings.ReadSettings()
	// routes.Init()

	route := gin.Default()
	route.GET("/get", func(ctx *gin.Context) {
		name := ctx.Query("name")
		if name == "Jon" {
			ctx.JSON(http.StatusOK,name)
			return
		}
		ctx.JSON(http.StatusBadRequest,"error")
		return
	})
	route.Run()
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
