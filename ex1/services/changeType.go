package services

import (
	"brt_adapter/models"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

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
func XmlType()(models.Paymen,error){

	var values models.Paymen
	
	req,err:=ioutil.ReadFile("./set.xml")
	if err!=nil{
		return values,errors.New("error with file read!!")
	}

	err=xml.Unmarshal([]byte(req),&values)
	if err!=nil{
		return values,errors.New("error with xml unmarshling!!")
	}
	return values,nil
}


// Function of object Person
func (p Person) String() string {
	return fmt.Sprintf("\t ID : %s - Name : %s - Starsign \n", p.Name, p.Starsign)
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