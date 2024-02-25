package main

import (
	//"brt_adapter/db"
	//"brt_adapter/models"
	//"brt_adapter/routes"
	"brt_adapter/routes"
	"brt_adapter/settings"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"fmt"
)


type Paymen struct{
	XMLName xml.Name `xml:"payment"`
	Persons []Person `xml:"person"`

}
type Person struct{
	XMLName xml.Name `xml:"person"`
	Name string `xml:"Name"`
	Starsign string `xml:"Starsign"`
}

type jsonPerson struct{
	Name string
	Starsign string
}
// Function of object Person
func (p Person) String() string {
	return fmt.Sprintf("\t ID : %s - Name : %s - Starsign \n", p.Name, p.Starsign)
}

func main() {
	// var (
	// 	str ="<payment><transID>123321</transID><result>Paymen accepted</result></payment><resultcode>200</resultcode><cardholder></cardholder><>"
	// )
	settings.AppSettings=settings.ReadSettings()

	routes.Init()

	// xmlFile, err := os.Open("set.xml")
	// if err != nil {
	// 	fmt.Println("Opening file error : ", err)
	// 	return
	// }
	// defer xmlFile.Close()

	// xmlData, _ := ioutil.ReadAll(xmlFile)

	// var c Paymen
	// xml.Unmarshal(xmlData, &c)

	// // Write XML on screen
	// fmt.Println(c.Persons)

	// // Convert to JSON
	// var person jsonPerson
	// var persons []jsonPerson

	// for _, value := range c.Persons {
	// 	person.Name = value.Name
	// 	person.Starsign = value.Starsign
	// 	persons = append(persons, person)
	// }

	// jsonData, err := json.Marshal(persons)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// // Write JSON on screen
	// fmt.Println(string(jsonData))

	// // Write to JSON file
	// jsonFile, err := os.Create("./Employees.json")

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer jsonFile.Close()

	// jsonFile.Write(jsonData)
	// jsonFile.Close()
	// fmt.Println(person.Name)
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
