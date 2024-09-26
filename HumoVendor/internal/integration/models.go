package integration

import (
	"time"
)

type PaymentRequestBody struct {
	ID                     int64  `json:"id"`
	Account                string `json:"account"`
	ReceiverAmount         string `json:"rec_amount"`
	ReceiverBilAccCurrency string `json:"rec_curr"`
	ProviderServiceID      string `json:"prov_service_id"`
}

type PaymentResponseBody struct {
	Status        StatusCode `json:"status"`
	StatusDetails string     `json:"status_details"`
	PaymentID     time.Time  `json:"payment_id"`
}

type StatusCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PostCheckRequestBody struct {
	ID        int64  `json:"id"`
	Account   string `json:"account"`
	ServiceID string `json:"serviceid"`
}

type PostCheckResponseBody struct {
	Status StatusCode `json:"status"`
}

type GetReceiverInfoRequestBody struct {
	Account           string `json:"account"`
	ProviderServiceID string `json:"prov_service_id"`
}

type GetReceiverInfoResponseBody struct {
	Status       StatusCode       `json:"status"`
	ReceiverInfo ReceiverInfoType `json:"rec_info"`
}

type ReceiverInfoType map[string]string
