package adapter

import (
	"time"
)

type Resp struct {
	Code        int      `json:"code"`
	Message     string   `json:"message"`
	Amount      string   `json:"amount"`
	Currency    string   `json:"currency"`
	Fx          string   `json:"fx"`
	ToPay       Topay    `json:"topay"`
	AccountInfo map[string]string `json:"accountinfo"`
}

type Topay struct {
	ID   string `json:"id"`
	Info string `json:"info"`
}


type ReqForPayment struct{
	Service string `json:"service"`
	UserID string `json:"userid"`
	Hash string `json:"hash"`
	Account string `json:"account"`
	Amount string `json:"amount"`
	Currency string `json:"currency"`
	TxnID string `json:"txnid"`
	Phone string `json:"phone"`
	Fee string `json:"fee"`
	ProviderID string `json:"providerid"`
	Last_Name string `json:"last_name"`
	First_Name string `json:"first_name"`
	Middle_Name string `json:"middle_name"`
	Sender_Birthday string `json:"sender_birthday"`
	ID_Series_Number string `json:"id_series_number"`
	Address string `json:"address"`
	Resident_City string `json:"resident_city"`
	Resident_Country string `json:"resident_country"`
	Postal_Code string `json:"postal_code"`
	Recipient_Name string `json:"recipient_name"`
}

type RespForPayment struct{
	ID          int    `json:"id"`
	DataTime    string `json:"datatime"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Status      string `json:"status"`
	StatusCode  int    `json:"statuscode"`
	Amount      string `json:"amount"`
	FX          string `json:"fx"`
	ToPay       any    `json:"topay"`
	AccountInfo any    `json:"accountinfo"`
}


type ReqForCheck struct {
	Account    string    `json:"account"`
	Amount     string    `json:"amount"`
	Currency   string    `json:"currency"`
	Hash       string    `json:"hash"`
	ProviderID int       `json:"providerid"`
	Service    string    `json:"service"`
	UserID     string    `json:"userid"`
	Datatime   time.Time `json:"datatime"`
}

type PreCheckResp struct {
	Message string `json:"message"`
}


type PostCheckRespError struct {
	ErrorCode int    `json:"error_code,omitempty"`
	ErrorText string `json:"error_text"`
	ErrorType string `json:"error_type"`
}
