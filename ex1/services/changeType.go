package services

import (
	"brt_adapter/models"
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// func TypeChanger(resp string) {
// 	req:=RemoveNonXMLCharacters(resp)

// }

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

var (
	str ="<transID>123321</transID><result>paymentaccepted</result>"
)

func XmlType(str string)(models.PaymentResponse,error){

	var values models.PaymentResponse
	if err:=xml.Unmarshal([]byte(str),&values); err!=nil{
		return values,errors.New("error with xml unmarshal!!")
	}
	return values,nil
}

func JsonType(str string)(models.PaymentResponse,error){

	var values models.PaymentResponse
	if err:=json.Unmarshal([]byte(str),values);err!=nil{
		return values,errors.New("error with json unmarshal!!")
	}
	return values,nil
}