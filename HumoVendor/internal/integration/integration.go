package integration

import (
	"alif/internal/adapter"
	"time"

	"fmt"
	"net/http"
	"strconv"
)

type AppSetings struct {
	alifUrl    string
	userID     string
	httpClient *http.Client
}

type PaymentStatusHumo int

type PaymentStatus int

// BRT response code
const (
	PaymentStatusPendingForGateway  PaymentStatus = 300 // Ожидает
	PaymentStatusSendedToGateway    PaymentStatus = 301 
	PaymentStatusProcessedByGateway PaymentStatus = 302
	PaymentStatusRejectedByGateway  PaymentStatus = 304
	PaymentStatusTimeoutedByGateway PaymentStatus = 303
)

const (
	PaymentStatusSuccessfully PaymentStatusHumo=202 // ПЛАТЕЖ ПРИНЯТ НА ОБРАБОТКУ
	PaymentStatusDublicate PaymentStatusHumo=208 // ПОВТОРНЫЙ ПЛАТЕЖ ИЛИ trxn_id РАНЕЕ ПРОВЕДЕН (ДУБЛИКАТ)
	PaymentStatusNotFound PaymentStatusHumo=428 // НЕ БЫЛ СОЗДАН createpayment ЛИБО ВРЕМЯ ЖИЗНИ ПРОШЛО
	PaymentStatusSignError PaymentStatusHumo=403 // НЕВЕРНАЯ ПОДПИСЬ(НЕУСПЕШНОЙ АВТОРИЗАЦИИ ЗАПРОСА)
	PaymentStatusStatusDetailsErr PaymentStatusHumo=400 // ОШИБКА ВЫПОЛНЕНИЯ С ОПИСАНИЕМ ОШИБКИ В ТЕГЕ
	PaymentStatusServiseOff PaymentStatusHumo=503 //СЕРВИС ОТКЛЮЧЕН ИЛИ УСЛУГА ЗАБЛОКИРОВАНА
	PaymentStatusBalance PaymentStatusHumo=507 //БАЛАНС ИСЧЕРПАН
	PaymentStatusRequestBody PaymentStatusHumo=500 //ОШИБКА ПРИ ВЫПОЛНЕНИИ ЗАПРОСА(НЕВЕРНЫЙ ФОРМАТ ЗАПРОСА)
)


type Adapter interface {
	PreCheck(req *adapter.ReqForCheck) (status int64, description string, rawInfo map[string]string, err error)
	Payment(req *adapter.Req) (status int64, description string, paymentID string, err error)
	PostCheck(req *adapter.Req) (status int64, description string, err error)
}

type Integration interface {
	ReceiverInfo(operID string, req *GetReceiverInfoRequestBody) *GetReceiverInfoResponseBody
	Payment(operID string, req *PaymentRequestBody) *PaymentResponseBody
	PostCheck(operID string, req *PostCheckRequestBody) *PostCheckResponseBody
}

type integration struct {
	adapter Adapter
}

func NewIntegration(adapter Adapter) Integration {
	return &integration{adapter: adapter}
}

func (i integration) ReceiverInfo(operID string, req *GetReceiverInfoRequestBody) *GetReceiverInfoResponseBody {
	var resp GetReceiverInfoResponseBody
	format := "Mon, 02 Jan 2006 15:04:05 -0700"
	datetime := time.Now().Format(format)
	reqBody := adapter.ReqForCheck{Account: req.Account, Service: req.ProviderServiceID, Datatime: datetime, Currency: "TJS"}

	status, desc, rawInfo, err := i.adapter.PreCheck(&reqBody)
	if err != nil {
		resp.Status.Code = int(FAILED)
		resp.Status.Message = err.Error()
		return &resp
	}

	switch status {
	case 200:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = desc
		resp.ReceiverInfo = rawInfo
	case 400, 402:
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = desc
		resp.ReceiverInfo = rawInfo
	case 500, 503:
		resp.Status.Code = int(PaymentStatusTimeoutedByGateway)
		resp.Status.Message = desc
		resp.ReceiverInfo = rawInfo
	case 521, 520, 409, 406:
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = desc
		resp.ReceiverInfo = rawInfo
	case 405, 404, 403:
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = desc
		resp.ReceiverInfo = rawInfo
	default:
		resp.Status.Code = int(PaymentStatusTimeoutedByGateway)
		resp.Status.Message = desc
		resp.ReceiverInfo = rawInfo
	}
	return &resp
}

func (i integration) Payment(operID string, req *PaymentRequestBody) *PaymentResponseBody {
	var resp PaymentResponseBody

	amount, err := strconv.ParseFloat(req.ReceiverAmount, 64)

	if err != nil {
		resp.Status.Code = int(FAILED)
		resp.Status.Message = err.Error()
		return &resp
	}
	reqBody := adapter.Req{Service: req.ProviderServiceID, TxnID: string(req.ID), Amount: amount, Currency: req.ReceiverBilAccCurrency, Account: req.Account}
	status, desc, payID, err := i.adapter.Payment(&reqBody)
	if err != nil {
		resp.Status.Code = int(FAILED)
		resp.Status.Message = err.Error()
		return &resp
	}
	fmt.Println(status, desc, payID)
	switch status {
	case SUCCESS:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = desc
		resp.ReceiverTrnID = payID
	
	case PENDING:
		resp.Status.Code = int(PaymentStatusPendingForGateway)
		resp.Status.Message = desc
		resp.ReceiverTrnID = payID
	
	case ACCEPTED:
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = desc
		resp.ReceiverTrnID = payID
	
	case FAILED:
		resp.Status.Code = int(PaymentStatusTimeoutedByGateway)
		resp.Status.Message = desc
		resp.ReceiverTrnID = payID
	
	case CANCELED:
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = desc
		resp.ReceiverTrnID = payID
	}
	return &resp
}

func (i integration) PostCheck(operID string, req *PostCheckRequestBody) *PostCheckResponseBody {
	var resp PostCheckResponseBody
	reqBody := adapter.Req{Account: req.Account, Service: req.ServiceID, TxnID: string(req.ID)}
	status, desc, err := i.adapter.PostCheck(&reqBody)
	fmt.Println(status, desc, err)
	if err != nil {
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = err.Error()
		return &resp
	}
	switch status {
	case SUCCESS:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = err.Error()
	
	case PENDING:
		resp.Status.Code = int(PaymentStatusPendingForGateway)
		resp.Status.Message = err.Error()
	
	case ACCEPTED:
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = err.Error()
		
	
	case FAILED:
		resp.Status.Code = int(PaymentStatusTimeoutedByGateway)
		resp.Status.Message = err.Error()
		
	
	case CANCELED:
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = err.Error()
		
	}
	return &resp
}
