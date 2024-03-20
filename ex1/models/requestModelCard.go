package models

import "encoding/xml"

type RequestCard struct {
	PAN        string
	ExpDate    string
	AgentLogin string
	Brand      string
}

type ResponceNPC struct {
	XMLName           xml.Name     `xml:"reply" json:"-"`
	Proto             string       `xml:"proto" json:"proto"`
	ResultCode        int          `xml:"resultcode" json:"resultcode"`
	Result            string       `xml:"result" json:"result"`
	Remain            float64      `xml:"remain" json:"remain"`
	Available         float64      `xml:"available" json:"available"`
	TranId            string       `xml:"tranid" json:"tranid"`
	Statement         []string     `xml:"list" json:"list"`
	CardStatement     *[]Statement `xml:"card_statement>operation"`
	Decline           string       `xml:"decline" json:"decline"`
	SessionId         string       `xml:"sessionid" json:"sessionid"`
	CardStatus        int          `xml:"cardstatus,omitempty" json:"cardstatus,omitempty"`
	CardStatusDesc    string       `xml:"cardstatusdesc,omitempty" json:"cardstatusdesc,omitempty"`
	AccountStatus     int          `xml:"accountstatus,omitempty" json:"accountstatus,omitempty"`
	AccountStatusDesc string       `xml:"accountstatusdesc,omitempty" json:"accountstatusdesc,omitempty"`
	CacheTime         string       `xml:"cacheTime,omitempty" json:"cache_time,omitempty"`
	Currency          string       `xml:"currency,omitempty" json:"currency,omitempty"`
	CardholderPhone   string       `xml:"CardholderPhone,omitempty" json:"CardholderPhone,omitempty"`
	ApprovalCode      string       `xml:"approval_code,omitempty" json:"approval_code,omitempty"`
}

type Statement struct {
	TranID         string `xml:"tran_id"`
	TransDate      string `xml:"trans_date"`
	OperationCode  string `xml:"operation_code"`
	OperationName  string `xml:"operation_name"`
	Amount         string `xml:"amount"`
	Currency       string `xml:"currency"`
	BankTerminalID string `xml:"bank_terminal_id"`
	TerminalPlace  string `xml:"terminal_place"`
	IsReversal     string `xml:"is_reversal"`
	AccAmount      amount `xml:"accAmount"`
	AccFee         amount `xml:"accFee"`
	TrnFee         amount `xml:"trnFee"`
}

type amount struct {
	Amount   string `xml:"amount"`
	Currency string `xml:"currency"`
}
