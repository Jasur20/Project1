package adapter

import "encoding/xml"

type PreCheckResp struct {
	Message string `json:"message"`
	//	Currency string `json:"currency,omitempty"`
}

type PaymentResp struct {
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
	ApprovalCode    string   `xml:"approval_code,omitempty"`
}

type PostCheckRespError struct {
	XMLName   xml.Name `xml:"responseError" json:"-"`
	ErrorCode int      `xml:"errorCode,omitempty" json:"error_code,omitempty"`
	ErrorText string   `xml:"errorText" json:"error_text"`
	ErrorType string   `xml:"errorType" json:"error_type"`
}
