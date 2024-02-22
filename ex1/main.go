package main

import (
	//"brt_adapter/db"
	//"brt_adapter/models"
	//"brt_adapter/routes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"io/ioutil"

	//"encoding/xml"
	"fmt"
	//"io/ioutil"

	//"brt_adapter/routes"
	"brt_adapter/settings"

	//"github.com/spf13/viper"
)

type Person struct{
	Name string
	Starsign string
}

func main() {
	settings.AppSettings=settings.ReadSettings()
	// fmt.Println(GetSha256(settings.AppSettings.Server,[]byte("2berXmPOpmjTsTLtrrpZBKgMt")))
	// fmt.Println(GetSha256(settings.AppSettings.TimeoutReq2NPCSec,[]byte("2berXmPOpmjTsTLtrrpZBKgMt")))
	// routes.Init()

	fileXml,err:=ioutil.ReadFile("./set.xml")
	if err!=nil{
		panic(err)
	}
	var c Person
	err=xml.Unmarshal([]byte(fileXml),&c)
	if err!=nil{
		panic(err)
	}
	fmt.Println(c.Name)


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
