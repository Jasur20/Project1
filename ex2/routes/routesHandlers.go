package routes

import (
	"brt_adapter/db"
	"brt_adapter/models"
	"brt_adapter/services"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func preCheck(ctx *gin.Context) {
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

func pay(ctx *gin.Context) {

	var errResp models.ErrorStruct

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
	// // APIURL := "http://"+"localhost"+":8080"+"/preCheckwithNum"
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

	//resString := fmt.Sprintf("<responseTransaction><transID>%s</transID><result>%s</result><resultCode>%s</resultCode><cardHolder>%s</cardHolder><cardNumberHash>%s</cardNumberHash><cardNumber>%s</cardNumber><reason>%s</reason><amount>%s</amount><paymenturl>%s</paymenturl><decline>%s</decline><approval_code>%s</approval_code></responseTransaction>",services.CreateTranId(),"Sending","200","Dastras",cardHash,cardNumber,"Отправка","2000","http://localhost:8080/Payment","Yes","300",)
	resString := fmt.Sprintf("<responseTransaction><transID>8sRhnaJP</transID><result>%s</result><resultCode>%s</resultCode><cardHolder>%s</cardHolder><cardNumberHash>%s</cardNumberHash><cardNumber>%s</cardNumber><reason>%s</reason><amount>%s</amount><paymenturl>%s</paymenturl><decline>%s</decline><approval_code>%s</approval_code></responseTransaction>","Sending","200","Dastras",cardHash,cardNumber,"Отправка","2000","http://localhost:8080/Payment","Yes","300",)

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
