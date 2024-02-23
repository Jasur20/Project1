package main

import (
	//"brt_adapter/db"
	//"brt_adapter/models"
	//"brt_adapter/routes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	//"encoding/json"
	"encoding/xml"

	//"io/ioutil"

	//"encoding/xml"
	"fmt"
	//"io/ioutil"

	//"brt_adapter/routes"
	"brt_adapter/models"
	"brt_adapter/services"
	"brt_adapter/settings"
	//"github.com/spf13/viper"
)

type Person struct{
	Name string
	Starsign string
}

func main() {
	var (
		str ="<payment><transID>123321</transID><result>Paymen accepted</result></payment><resultcode>200</resultcode><cardholder></cardholder><>"
	)
	settings.AppSettings=settings.ReadSettings()
	// fmt.Println(GetSha256(settings.AppSettings.Server,[]byte("2berXmPOpmjTsTLtrrpZBKgMt")))
	// fmt.Println(GetSha256(settings.AppSettings.TimeoutReq2NPCSec,[]byte("2berXmPOpmjTsTLtrrpZBKgMt")))
	// routes.Init()
	// fileXml,err:=ioutil.ReadFile("./set.xml")
	// if err!=nil{
	// 	panic(err)
	// }
	// fmt.Println(string(fileXml))
	req:=services.RemoveNonXMLCharacters(str)
	var f models.Ex
	err:=xml.Unmarshal([]byte(req),&f)
	if err!=nil{
		panic(err)
	}
	fmt.Println(f.TrancID)
	fmt.Println(f.Result)
	
	// values,err:=services.JsonType(req)
	// if err!=nil{
	// 	panic(err)
	// }
	// fmt.Println(values.Amount)


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
