package usecase

type ReqForPreCheck struct {
}

type ReqForCheck struct {
	Command   string  `json:"command"`
	Login    string  `json:"login"`
	Account  string  `json:"account"`
	Txn_ID   string  `json:"txnid"`
	Prv_ID   string  `json:"prvid"`
	Sum      float64 `json:"sum"`
	Sign     string  `json:"sign"`
	Txn_Date string  `json:"txndate"`
	Ccy      string  `json:"ccy"`
}

type GetReceiverInfoResponseBody struct {
	Status       StatusCode        `json:"status"`
	ReceiverInfo map[string]string `json:"rec_info"`
}

type StatusCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PostCheckResponseBody struct {
	Status StatusCode `json:"status"`
}

type GetReceiverInfoRequestBody struct {
	Account           string `json:"account"`
	ProviderServiceID string `json:"prov_service_id"`
}

type PaymentRequestBody struct {
	ID                     int64  `json:"id"`
	Account                string `json:"account"`
	ReceiverAmount         string `json:"rec_amount"`
	ReceiverBilAccCurrency string `json:"rec_curr"`
	ProviderServiceID      string `json:"prov_service_id"`
}

type PostCheckRequestBody struct {
	ID        int64  `json:"id"`
	ServiceID string `json:"serviceid"`
}

type ReqForPayment struct {
	Command    string  `yaml:"command"`
	Login     string  `yaml:"login"`
	Account   string  `yaml:"account"`
	Ccy       string  `yaml:"ccy"`
	Txn_ID    string  `yaml:"txnid"`
	Prv_ID    string  `yaml:"prvid"`
	Sum       float64 `yaml:"sum"`
	Sign      string  `yaml:"sign"`
	Cr_Amount string  `yaml:"cramount"`
	Txn_Date  string  `yaml:"txndate"`
	Wallet    string  `yaml:"wallet"`
}

type PaymentResponseBody struct {
	Status        StatusCode `json:"status"`
	StatusDetails string     `json:"status_details"`
	PaymentID     string     `json:"payment_id"`
}

type ReqForPostCheck struct {
	Command string `yaml:"command"`	
	Login  string `yaml:"Login"`
	TxnId  string `yaml:"txn_id"`
	Sign   string `yaml:"Sign"`
}
