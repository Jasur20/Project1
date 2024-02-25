package services

import "encoding/xml"

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