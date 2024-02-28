package models

import "encoding/xml"

type PostCheckResp struct {
	XMLName       xml.Name `xml:"response"`
	Status        int      `xml:"status"`
	StatusDetails string   `xml:"statusdetails"`
	ReceiptID     string   `xml:"receipt_id"`
}
