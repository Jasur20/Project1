package routes

import (
	"brt_adapter/db"
	"brt_adapter/models"
	"brt_adapter/services"
	"brt_adapter/settings"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"net/url"

	"github.com/gin-gonic/gin"
)

var (
	info = "/info"
)

var response services.Paymen

func preCheckwithNum(ctx *gin.Context) {

	var list models.List
	var errResp models.ErrorStruct

	if err := ctx.ShouldBindJSON(&list); err != nil {
		errResp.Status = 400
		errResp.StatusDetails = "Bad type input"
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	err := services.ValidationNum(list.Phone)
	if err != nil {
		errResp.Status = 400
		errResp.StatusDetails = "error in verification Phone!!"
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	list = models.List{Phone: list.Phone, PAN: list.PAN, Data: list.Data, CVV: list.CVV, NAME: list.NAME, Status: list.Status}
	rows, status, err := db.FindCardsBankWithNum(list.Phone)
	if err != nil {
		errResp.Status = status
		errResp.StatusDetails = err.Error()
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	ctx.JSON(http.StatusOK, rows)
}

func preCheck(ctx *gin.Context) {

	var errResp models.ErrorStruct
	var values models.BankCard
	//var response models.PreCheckResp

	if err := ctx.ShouldBindJSON(&values); err != nil {
		errResp.Status = 400
		errResp.StatusDetails = "Bad type input "
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	err := services.ValidationPan(values.PAN)
	if err != nil {
		errResp.Status = 400
		errResp.StatusDetails = err.Error()
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	values = models.BankCard{PAN: values.PAN, Data: values.Data, CVV: values.CVV, NAME: values.NAME, Status: values.Status}

	bankcardtype := services.BankCardVerification(values.PAN)
	if bankcardtype != "BRT card" {
		// rows, status, err := services.SendRequestToFimi(values.PAN)
		defer func(ctx *gin.Context) {
			var request models.RequestCard

			request.PAN = ctx.Query("pan")
			request.ExpDate = ctx.Query("exp")
			request.AgentLogin = ctx.Query("agentlogin")
			request.Brand = ctx.Query("brand")

			// APIURL := "http://"+"localhost"+":8080"+"/getCardInfoKM"
			// client := http.Client{Timeout: settings.AppSettings.AppParams.TimeoutReq2NPCSec * time.Second}
			// url:="http://ya.ru"

			// resp, err := client.Post(url, "application/xml", nil)
			// if err!=nil{
			// 	errResp.Status=404
			// 	errResp.StatusDetails=err.Error()
			// 	ctx.JSON(http.StatusBadRequest,errResp)
			// 	return
			// }
			// defer resp.Body.Close()
			// resString,err:=ioutil.ReadAll(resp.Body)
			// if err!=nil{
			// 	errResp.Status=400
			// 	errResp.StatusDetails=err.Error()
			// 	ctx.JSON(http.StatusBadRequest,errResp)
			// 	return
			// }

		}()

	}
	fmt.Println("BRT card")
	rows, status, err := db.FindCardsBankWithPan(values.PAN)
	if err != nil {
		errResp.Status = status
		errResp.StatusDetails = err.Error()
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	ctx.JSON(http.StatusOK, rows)

}

type ReqISPC struct {
	AgentLogin   string
	AgentPassw   string
	CardHash     string
	CardNumber   string
	ClientCode   string
	AgentTransId string
	Reason       string
	Amount       string
}

func pay(ctx *gin.Context) {

	var errResp models.ErrorStruct
	var req ReqISPC
	agentLogin := ctx.PostForm("agentLogin")
	agentPassw := ctx.PostForm("agentPassword")
	cardHash := ctx.PostForm("cardHash")
	cardNumber := ctx.PostForm("cardNumber")
	clientCode := ctx.PostForm("clientCode")
	agentTransID := ctx.PostForm("transID")
	reason := ctx.PostForm("reason")
	amount, err := strconv.ParseFloat(ctx.PostForm("amount"), 64)
	if err != nil {
		errResp.Status = 400
		errResp.StatusDetails = err.Error()
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	
	fmt.Println(agentLogin, agentPassw, cardHash, cardNumber, clientCode, agentTransID, reason, amount)
	APIURL := "http://" + "localhost" + ":8080" + "/v2cTransfer"
	client := http.Client{Timeout: settings.AppSettings.AppParams.TimeoutReq2NPCSec * time.Second}
	//url:="http://ya.ru"
    body := url.Values{}
	body.Add("agentlogin",agentLogin)
	
	resp, err := client.Post(APIURL, "application/x-www-form-urlencoded", )
	if err != nil {
		errResp.Status = 404
		errResp.StatusDetails = err.Error()
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	defer resp.Body.Close()
	resString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errResp.Status = 400
		errResp.StatusDetails = err.Error()
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	resString := "<responseTransaction><transID>qwe1234g</transID><result>Send</result><resultCode>200</resultCode><cardHolder>Dastras</cardHolder><cardNumberHash>1234432112344321</cardNumberHash><cardNumber>5058271232121223</cardNumber><reason>BRT</reason><amount>1420</amount><paymenturl>localhost:8080</paymenturl><decline>NO</decline><approval_code>200</approval_code></responseTransaction>"

	var values models.PaymentResponse
	err = xml.Unmarshal([]byte(resString), &values)
	if err != nil {
		panic(err)
	}
	err = db.SavePay(values.TransID, values.Result, values.ResultCode, values.CardHolder, values.CardNumberHash, values.MaskedPanNumber, values.Reason, values.Amount, values.PaymentURL, values.Decline, values.ApprovalCode)
	if err != nil {
		errResp.Status = 400
		errResp.StatusDetails = err.Error()
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	ctx.JSON(http.StatusOK, values)

}

func postCheck(ctx *gin.Context) {

	var errRsp models.ErrorStruct
	tranID := ctx.PostForm("trnx_id")

	resp, err := db.PostCheck(tranID)
	if err != nil {
		errRsp.Status = 400
		errRsp.StatusDetails = err.Error()
		ctx.JSON(http.StatusBadRequest, errRsp)
		return
	}
	ctx.JSON(200, resp)
}
