package integration

import (
	"fmt"
	"strconv"
)

type Adapter interface {
	PreCheck(account string, serviceID string) (status int64, description string, rawInfo map[string]string, err error)
	//Payment(account, amount, trnID, notifyRoute string) (status int64, description string, paymentID int64, err error)
	Payment(account, serviceID, amount, trnID, notifyRoute string) (status int64, description string, paymentID string, err error)
	PostCheck(trnID string) (status int64, description string, err error)
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

type PaymentStatus int

const (
	PaymentStatusPendingForGateway  PaymentStatus = 300 // Ожидает
	PaymentStatusSendedToGateway    PaymentStatus = 301
	PaymentStatusProcessedByGateway PaymentStatus = 302
	PaymentStatusRejectedByGateway  PaymentStatus = 304
	PaymentStatusTimeoutedByGateway PaymentStatus = 303
)

func (i integration) ReceiverInfo(operID string, req *GetReceiverInfoRequestBody) *GetReceiverInfoResponseBody {
	var resp GetReceiverInfoResponseBody
	status, desc, rawInfo, err := i.adapter.PreCheck(req.Account, req.ProviderServiceID)
	if err != nil {
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = err.Error()
		return &resp
	}

	switch status {
	case 200:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = desc
		resp.ReceiverInfo = rawInfo
	default:
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = desc
		resp.ReceiverInfo = rawInfo
	}
	return &resp
}

func (i integration) Payment(operID string, req *PaymentRequestBody) *PaymentResponseBody {
	var resp PaymentResponseBody

	amount, err := strconv.ParseFloat(req.ReceiverAmount, 64)

	if err != nil {
		resp.Status.Code = int(PaymentStatusTimeoutedByGateway)
		resp.Status.Message = err.Error()
		return &resp
	}

	status, desc, payID, err := i.adapter.Payment(req.Account, req.ProviderServiceID, fmt.Sprintf("%.2f", amount), fmt.Sprint(req.ID), "")
	if err != nil {
		resp.Status.Code = int(PaymentStatusTimeoutedByGateway)
		resp.Status.Message = err.Error()
		return &resp
	}
	switch status {
	case 310:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = desc
		resp.ReceiverTrnID = payID
	case 320:
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = desc
	default:
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = desc
		resp.ReceiverTrnID = payID
	}
	return &resp
}

func (i integration) PostCheck(operID string, req *PostCheckRequestBody) *PostCheckResponseBody {
	var resp PostCheckResponseBody
	status, desc, err := i.adapter.PostCheck(fmt.Sprint(req.ID))
	if err != nil {
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = err.Error()
		return &resp
	}
	switch status {
	case 310:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = desc
	case 320, 404:
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = desc
	default:
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = desc
	}
	return &resp
}
