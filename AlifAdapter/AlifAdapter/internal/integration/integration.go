package integration

import (
	"alif/internal/adapter"
	"fmt"
	"strconv"
)

type PaymentStatusAlif int

const(

)

type PaymentStatus int

//BRT response code
const (
	PaymentStatusPendingForGateway  PaymentStatus = 300 // Ожидает
	PaymentStatusSendedToGateway    PaymentStatus = 301
	PaymentStatusProcessedByGateway PaymentStatus = 302
	PaymentStatusRejectedByGateway  PaymentStatus = 304
	PaymentStatusTimeoutedByGateway PaymentStatus = 303
)

//Alif responce code
const (
	PaymentStatusSuccessfully               PaymentStatusAlif = 200 // успешно
	PaymentStatusConversionError            PaymentStatusAlif = 285 // произошла оошибка при конвертации
	PaymentStatusCourseCurrencyHasChanged   PaymentStatusAlif = 286 // курс валют изменился
	PaymentStatusInvalidRequest             PaymentStatusAlif = 400 // неверный запрос
	PaymentStatusNotAuthorized              PaymentStatusAlif = 401 // не авторизован
	PaymentStatusRecepienWasNotFound        PaymentStatusAlif = 402 // Получатель не найден
	PaymentStatusNoAccess                   PaymentStatusAlif = 403 // нет доступа
	PaymentStatusPaymentNotFound            PaymentStatusAlif = 404 // Платеж не найден
	PaymentStatusMethodNotAllowed           PaymentStatusAlif = 405 // Метод не разрешен
	PaymentStatusReConfirmationPayment      PaymentStatusAlif = 406 // Повторное подтверждение платежа
	PaymentStatusRepeatedVerificationReques PaymentStatusAlif = 409 // Повторный запрос проверки
	PaymentStatusInvalidAccount             PaymentStatusAlif = 410 // неверный аккаунт получателя
	PaymentStatusAmountSmall                PaymentStatusAlif = 411 // Сумма слишком мала
	PaymentStatusAmountLarge                PaymentStatusAlif = 412 // Сумма слишком велика
	PaymentStatusIncorrectTransferAmount    PaymentStatusAlif = 413 // Неверная сумма перевода
	PaymentStatusInvalidRequestIdentidier   PaymentStatusAlif = 414 // Неверный идентификатор запроса
	PaymentStatusClientStopList             PaymentStatusAlif = 415 // Клиент в стоп-листе
	PaymentStatusInternalServerError        PaymentStatusAlif = 500 // Внутренняя ошибка сервера
	PaymentStatusTemporaryErrorRepeatLater  PaymentStatusAlif = 503 // Временная ошибка. Повторите запрос позже
	PaymentStatusPaymentPending             PaymentStatusAlif = 520 // Платеж в ожидании
	PaymentStatusPaymentChecking            PaymentStatusAlif = 521 // Платеж на проверке
)

//Alif StatusCode
const(
	ACCEPTED =0
	SUCCESS=1
	PENDING=2
	FAILED=3
	CANCELED=4
)


type Adapter interface {
	PreCheck(data *adapter.CheckAccountReq) (status int64, description string, rawInfo map[string]string, err error)
	//Payment(account, amount, trnID, notifyRoute string) (status int64, description string, paymentID int64, err error)
	Payment(data *adapter.Req) (status int64, description string, paymentID string, err error)
	PostCheck(data *adapter.Req) (status int64, description string, err error)
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
	data:=adapter.CheckAccountReq{Service:req.ProviderServiceID,Account: req.Account,Currency: "TJS"}
	status, desc, rawInfo, err := i.adapter.PreCheck(&data)
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
	case 400,402:
		resp.Status.Code=int(PaymentStatusRejectedByGateway)
		resp.Status.Message=desc
		resp.ReceiverInfo=rawInfo
	case 500:
		resp.Status.Code=int(PaymentStatusTimeoutedByGateway)
		resp.Status.Message=desc
		resp.ReceiverInfo=rawInfo
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
	data:=adapter.Req{Service:req.ProviderServiceID,Account: req.Account,Amount:amount,Currency: req.ReceiverBilAccCurrency,TxnID:string(req.ID)}

	status, desc, payID, err := i.adapter.Payment(&data)
	if err != nil {
		resp.Status.Code = int(FAILED)
		resp.Status.Message = err.Error()
		return &resp
	}
	fmt.Println(status,desc,payID)
	switch status {
	case 200:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = desc
		resp.ReceiverTrnID = payID
	case 411,412,413:
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = desc
		resp.ReceiverTrnID=payID
	case 500,503,520:
		resp.Status.Code=int(PaymentStatusTimeoutedByGateway)
		resp.Status.Message=desc
		resp.ReceiverTrnID=payID
	}
	return &resp
}


func (i integration) PostCheck(operID string, req *PostCheckRequestBody) *PostCheckResponseBody {
	var resp PostCheckResponseBody

	data:=adapter.Req{Account:req.Account,Service:req.ServiceID, TxnID:string(req.ID),Currency:"TJS"}

	status, desc, err := i.adapter.PostCheck(&data)
	fmt.Println(status,desc,err)
	if err != nil {
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = err.Error()
		return &resp
	}
	switch status {
	case 200:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = desc
	case 500,503:
		resp.Status.Code = int(PaymentStatusTimeoutedByGateway)
		resp.Status.Message = desc
	default:
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = desc
	}
	return &resp
}

