package models

import "encoding/xml"

type PreCheckResp struct {
	XMLName       xml.Name `xml:"response"`
	Status        int      `xml:"status"`
	StatusDetails string   `xml:"statusdetails"`
	PreCheckInfo  struct {
		RawInfo string `xml:"RawInfo"`
		Items   struct {
			Name     string `xml:"name,omitempty"`
			Address  string `xml:"address,omitempty"`
			Previous string `xml:"previous,omitempty"`
			Present  string `xml:"present,omitempty"`
			Date     string `xml:"date,omitempty"`
			Rest     string `xml:"rest,omitempty"`
			Item     []struct {
				Label string `xml:"label,attr"`
				Value string `xml:"value,attr"`
			} `xml:"item"`
		} `xml:"items"`
	} `xml:"precheckinfo"`
}

type ErrorStruct struct {
	ErrorName     string   
	Status        int      
	StatusDetails string  
}
