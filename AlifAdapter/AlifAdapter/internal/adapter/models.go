package adapter

type Resp struct{
	ID string `json:"id"`
	DataTime string `json:"datatime"`
	Code int `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
	StatusCode int `json:"statusCode"`
	Amount string `json:"amount"`
	Fx string `json:"fx"`
	ToPay interface{} `json:"topay"`
	AccountInfo string `json:"accountinfo"`
}

type Req struct{
	Service string `json:"service"`
	UserID string `json:"userid"`
	Hash string `json:"hash"`
	Account string `json:"account"`
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	TxnID string `json:"txnid"`
	Phone string `json:"phone"`
	Fee string `json:"fee"`
	Providerid string `json:"proviredid"`
	Last_Name string `json:"last_name"`
	First_Name string `json:"first_name"`
	Middle_Name string `json:"middle_name"`
	Sender_Birthday string`json:"sender_birhtday"`
	Address string `json:"address"`
	Resident_Country string `json:"resident_country"`
	Postal_Code string `json:"postal_code"`
	Recipient_Name string`json:"recipient_name"`
}

type CheckAccountReq struct{
	Service string `json:"service"`
	UserID string `json:"userid"`
	Hash string `json:"hash"`
	Account string `json:"account"`
	Currency string `json:"currency"`
	ProviderID string `json:"providerid"`
	Datatime string `json:"datatime"`
}

type PreCheckResp struct {
	Message string `json:"message"`
	//	Currency string `json:"currency,omitempty"`
}

type CheckAccountResp struct{
	Code string `json:"code"`
	Message string `json:"message"`
	AccountInfo string `json:"accountinfo"`
	ToPay string `json:"topay"`
	Amount string `json:"amount"`
	Currency string `json:"currency"`
	Fx string `json:"fx"`
}

type PaymentRes struct{
	ID int `json:"id"`
	DataTime string `json:"datatime"`
	Code int `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
	StatusCode int `json:"statuscode"`
	Amount string `json:"amount"`
	FX string `json:"fx"`
	ToPay any `json:"topay"`
	AccountInfo any `json:"accountinfo"`
}

type PostCheckRespError struct {
	ErrorCode int       `json:"error_code,omitempty"`
	ErrorText string   `json:"error_text"`
	ErrorType string   `json:"error_type"`
}

