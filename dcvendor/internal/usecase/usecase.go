package usecase

import (
	"fmt"
	"strconv"
	"time"
)

type PaymentStatusDC int

type PaymentStatus int

// BRT response code for Paymont
const (
	PaymentStatusPendingForGateway  PaymentStatus = 300 // Ожидает
	PaymentStatusSendedToGateway    PaymentStatus = 301 // Принял шлюз DC
	PaymentStatusProcessedByGateway PaymentStatus = 302 // Успешно прошел Dc
	PaymentStatusRejectedByGateway  PaymentStatus = 304 // Отмена
	PaymentStatusTimeoutedByGateway PaymentStatus = 303 // Превышение времени ожидания
)

// DC response code
const (
	PaymentStatusSuccessfully              PaymentStatusDC = 0   // OK
	PaymenStatusInvalidIdentifier          PaymentStatusDC = 4   // Неверный формат идентификатора
	PaymentStatusIDNotFount                PaymentStatusDC = 5   // Идентификатор абонента не найден
	PaymentStatusErrSign                   PaymentStatusDC = 13  // Ошибка подписи
	PaymentStatusPayNotFinish              PaymentStatusDC = 90  // Проведение платежа не окончено
	PaymentStatusNeedActivatecard          PaymentStatusDC = 97  // Необходимо активировать карту
	PaymentStatusPayRejected               PaymentStatusDC = 12  // Платеж отменен. Срества обратно возвращены на ваш счет.
	PaymentStatusLackOfFunds               PaymentStatusDC = 220 // Нехватка средств на вешм счете
	PaymentStatusAmountSmall               PaymentStatusDC = 241 // Сумма слишком мала
	PaymentStatusAmountBig                 PaymentStatusDC = 242 // Сумма слишко велика
	PaymentStatusAnotherMistake            PaymentStatusDC = 300 // Другая ошибка поставщика услуг
	PaymentStatusTrnxNotFound              PaymentStatusDC = 6   // Другая ошибка поставщика услуг
	PaymentStatusRepeatReqLater            PaymentStatusDC = 1   // Переменная ошибка. Повторите запрос позже
	PaymentStatusPayAcceptanceBySupplier   PaymentStatusDC = 7   // Прием платежа запрещен поставщиком
	PaymentStatusPayAcceptanceTechReasoncs PaymentStatusDC = 8   // Причем платежа по техническим причинам
	PaymentStatusReqErr                    PaymentStatusDC = 10  // Request error
)

type Adapter interface {
	PreCheck(req *ReqForCheck) (status int, description string, rawInfo map[string]string, err error)
	Payment(req *ReqForPayment) (status int, description string, paymentID string, err error)
	PostCheck(req *ReqForPostCheck) (status int, description string, err error)
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

	format := "20060102150405"
	datetime := time.Now().Format(format)

	reqBody := ReqForCheck{Command: "check", Account: req.Account, Prv_ID: req.ProviderServiceID,Sum:10.45,Ccy: "TJS", Txn_Date: datetime}
	status, desc, rawInfo, err := i.adapter.PreCheck(&reqBody)
	if err != nil {
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = err.Error()
		return &resp
	}

	switch status {
	case 0:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = desc
		resp.ReceiverInfo = rawInfo
	case 4, 5, 13, 241, 242:
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

	format := "20060102150405"
	datetime := time.Now().Format(format)
	amount, err := strconv.ParseFloat(req.ReceiverAmount, 64)

	if err != nil {
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = err.Error()
		return &resp
	}
	reqBody := ReqForPayment{Command:"pay",Account:req.Account,Ccy: req.ReceiverBilAccCurrency, Txn_ID: fmt.Sprint(req.ID), Prv_ID: req.ProviderServiceID,Sum: amount,Txn_Date: datetime}
	status, desc, payID, err := i.adapter.Payment(&reqBody)
	if err != nil {
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = err.Error()
		return &resp
	}

	switch status {
	case 0:
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = desc
		resp.PaymentID = payID
	case 90,97,12,13,241,242,300,1,7,8,10,4:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = desc
		resp.PaymentID = payID
	default:
		resp.Status.Code = int(PaymentStatusTimeoutedByGateway)
		resp.Status.Message = desc
		resp.PaymentID = payID
	}
	return &resp
}

func (i integration) PostCheck(operID string, req *PostCheckRequestBody) *PostCheckResponseBody {
	var resp PostCheckResponseBody
	reqBody := ReqForPostCheck{Command:"getstatus", TxnId:fmt.Sprint(req.ID)}
	status, desc, err := i.adapter.PostCheck(&reqBody)

	if err != nil {
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = err.Error()
		return &resp
	}
	switch status {
	case 0:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = desc
	case 13,90,4,12,220,300,6,97:
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = desc
	default:
		resp.Status.Code = int(PaymentStatusTimeoutedByGateway)
		resp.Status.Message = desc
	}
	return &resp
}
