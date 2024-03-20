package services

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	//"io/ioutil"
	"strings"
)

// delete unnecessary details
func RemoveNonXMLCharacters(xmlSTR string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 0x13, 0x10, 0x09:
			return -1
		default:
			return r
		}
	}, xmlSTR)
}

//open xml file and send option
func XmlType()(Paymen,error){

	var values Paymen
	
	req,err:=ioutil.ReadFile("./set.xml")
	if err!=nil{
		return values,errors.New("error with file read!!")
	}

	err=xml.Unmarshal([]byte(req),&values)
	if err!=nil{
		return values,errors.New("error with xml unmarshling!!")
	}
	fmt.Println(values.Persons)
	return values,nil
}


// Function of object Person
func (p Person) String() string {
	return fmt.Sprintf("\t ID : %s - Name : %s - Starsign :%d -Amount: %s-Data \n", p.Name, p.Starsign,p.Amount,p.Data)
}

// Convert xml option to json type
func JsonType()(jsonPerson,error){

	// open xml file
	xmlFile, err := os.Open("set.xml")
	if err != nil {
		fmt.Println("Opening file error : ", err)
		return jsonPerson{},errors.New("error with file open!!")
	}
	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	var c Paymen
	xml.Unmarshal(xmlData, &c)

	// Convert to JSON
	var person jsonPerson
	var persons []jsonPerson

	for _, value := range c.Persons {
		person.Name = value.Name
		person.Starsign = value.Starsign
		person.Amount=value.Amount
		person.Data=value.Data
		persons = append(persons, person)
	}

	jsonData, err := json.Marshal(persons)

	if err != nil {
		os.Exit(1)
		return jsonPerson{},errors.New("error with Marshal!!")
	}

	// Write to JSON file
	jsonFile, err := os.Create("./Employees.json")

	if err != nil {
		return jsonPerson{},errors.New("error with create file for settings!!")
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	defer jsonFile.Close()

	return person,nil
}

func CreateTranId() string{

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
    "abcdefghijklmnopqrstuvwxyzåäö" +
    "0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
    b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // Например "ExcbsVQs"
	return str
}
