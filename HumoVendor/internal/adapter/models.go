package adapter



type RespForPreCheck struct {
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


type Req struct{
	Request string `json:"request"`
	PS_ID string `json:"ps_id"`
	Amount float64 `json:"amount"`
	Account string `json:"account"`
	Service_ID string `json:"service_id"`
	Currency string `json:"currency"`
	Notify_Flag string `json:"notify_flag"`
	Trnx_ID float64 `json:"trnx_id"`
	Agent_Term_Receipt string `json:"agent_term_receipt"`
	Notify_Route string `json:"notify_route"`
	Sinn string `json:"sinn"`
	Hash string `json:"hash"`
	Date string `json:"date"`
}

type Resp struct{
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
	Datatime   string `json:"datatime"`
}
