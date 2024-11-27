package adapter

type PreCheckResp struct {
	OsmpTxnId    string  `xml:"osmp_txn_id"`
	PrvTxn       string  `xml:"prv_txn"`
	Sum          float64 `xml:"sum"`
	Ccy          string  `xml:"ccy"`
	Result       int     `xml:"result"`
	Comment      string  `xml:"comment"`
	Txn_Date     int64   `xml:"txn_date"`
	CreditAmount float64 `xml:"credit_amount"`
	CreditCurr   string  `xml:"credit_curr"`
	CurrRate     float64 `xml:"curr_rate"`
	Comis        float64 `xml:"comis"`
	Info         string  `xml:"info"`
}

type PaymentResp struct {
	OsmpTxnID string  `xml:"osmp_txn_id"`
	PrvTxn    string  `xml:"prv_txn"`
	Sum       float64 `xml:"sum"`
	Ccy       string  `xml:"ccy"`
	CrAmount  string  `xml:"cr_amount"`
	Result    int     `xml:"result"`
	Comment   string  `xml:"comment"`
	TxnDate   int64     `xml:"txn_date"`
}

type PostCheckResp struct {
	OsmpTxnId string  `xml:"osmp_txn_id"`
	Prv_Txn   string  `xml:"prv_txn"`
	Sum       float64 `xml:"Sum"`
	Ccy       string  `xml:"Ccy"`
	Result    int     `xml:"Result"`
	Comment   string  `xml:"comment"`
}
