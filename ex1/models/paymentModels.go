package models

import "encoding/xml"

type PaymentList struct{
	XMLName xml.Name `xml:"payment"`
	PaymentResponses []PaymentResponse `xml:"paymentresponse"`
}
type PaymentResponse struct {
	XMLName         xml.Name `xml:"responseTransaction"`
	TransID         string   `xml:"transID"`
	Result          string   `xml:"result"`
	ResultCode      int      `xml:"resultCode"`
	CardHolder      string   `xml:"cardHolder,omitempty"`
	CardNumberHash  string   `xml:"cardNumberHash,omitempty"`
	MaskedPanNumber string   `xml:"cardNumber,omitempty"`
	Reason          string   `xml:"reason"`
	Amount          float64  `xml:"amount"`
	PaymentURL      string   `xml:"paymenturl,omitempty"`
	Decline         string   `xml:"decline,omitempty"`
	ApprovalCode    string   `xml:"approval_code,omitempty"` // ispc4pos
}
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
