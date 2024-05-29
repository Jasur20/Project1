package integration

import (
	"alif/internal/adapter"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
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


// type integration struct{
	


type Adapter interface {
	PreCheck(body []byte) (status int64, description string, rawInfo map[string]string, err error)
	//Payment(account, amount, trnID, notifyRoute string) (status int64, description string, paymentID int64, err error)
	Payment(account, serviceID, amount, trnID, notifyRoute string) (status int64, description string, paymentID string, err error)
	PostCheck(trnID string,account string,serviceID string) (status int64, description string, err error)
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



func GetSha256(text string, secret []byte) string {

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(text))
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}


func (i integration) ReceiverInfo(operID string, req *GetReceiverInfoRequestBody) *GetReceiverInfoResponseBody {
	var resp GetReceiverInfoResponseBody
	body,err:=SendRequest(req.Account,req.ProviderServiceID,"0","0","0")
	if err!=nil{
		resp.Status.Code=(FAILED)
		resp.Status.Message=err.Error()
		return &resp
	}
	status, desc, rawInfo, err := i.adapter.PreCheck(body/*req.Account, req.ProviderServiceID*/)
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

	status, desc, payID, err := i.adapter.Payment(req.Account, req.ProviderServiceID, fmt.Sprintf("%.2f", amount), fmt.Sprint(req.ID), "")
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
	status, desc, err:= i.adapter.PostCheck(fmt.Sprint(req.ID),req.Account,req.ServiceID)
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


func SendRequest(account,serviceid,amount,trnID,notifyRoute string,req any )(req *http.Request,err error){
	var a adapter
	var valueForCheckType adapter.ReqForCheck
	if reflect.TypeOf(req)!=reflect.TypeOf(valueForCheckType){
		time:=time.Now()
		hash:=GetSha256(a.userid+account+"0"+"0",[]byte(a.userid))
    	data,err:=json.Marshal(adapter.ReqForPayment{Account: account,Currency: "TJS",Hash: hash,ProviderID: 0,Service: serviceID,UserID: a.userid,Datatime:time})
		if err!=nil{
			logrus.WithError(err).Warn("on marshaling")
			return nil,err
		}
		req,err=http.NewRequest(http.MethodPost,a.alifUrl,bytes.NewBuffer(data))
		if err!=nil{
			logrus.WithError(err).Info("on req")
			return nil,nil	
		}
		req.Header.Add("Content-Type", "application/json; charset=utf8")
		return req,nil
	}
	time:=time.Now()
	hash:=GetSha256(a.userid+account+"0"+"0",[]byte(a.userid))
    data,err:=json.Marshal(adapter.ReqForPayment{Account: account,Currency: "TJS",Hash: hash,ProviderID: 0,Service: serviceID,UserID: a.userid,Datatime:time})
	if err!=nil{
		logrus.WithError(err).Warn("on marshaling")
		return nil,err
	}
	req,err=http.NewRequest(http.MethodPost,a.alifUrl,bytes.NewBuffer(data))
	if err!=nil{
		logrus.WithError(err).Info("on req")
		return nil,nil	
	}
	req.Header.Add("Content-Type", "application/json; charset=utf8")
	return req,nil

	// res,err:=a.httpClient.Do(req)
	// if err!=nil{
	// 	logrus.WithError(err).Info("on response")
	// 	return nil,err
	// }
	// defer res.Body.Close()
	// body,err=io.ReadAll(res.Body)
	// if err!=nil{
	// 	logrus.WithError(err).Info("on reading responce body")
	// 	return nil,err
	// }
	// logrus.WithFields(
	// 	logrus.Fields{
	// 		"body": string(body),
	// 		"status": res.StatusCode,
	// 	},
	// ).Info("response to adapter")
	// return body,nil
}