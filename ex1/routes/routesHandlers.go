package routes

import (
	"brt_adapter/db"
	"brt_adapter/models"
	"brt_adapter/services"
	"brt_adapter/settings"
	//"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	//"github.com/spf13/viper"
)

var (
	info="/info"
)

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

	list = models.List{Phone: list.Phone,PAN: list.PAN,Data: list.Data,CVV: list.CVV,NAME: list.NAME,Status: list.Status}
	rows, status, err := db.FindCardsBankWithNum(list.Phone)
	if err != nil {
		errResp.Status = status
		errResp.StatusDetails = err.Error()
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	ctx.JSON(http.StatusOK, rows)
}

func preCheckwithPan(ctx *gin.Context) {

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

	values = models.BankCard{PAN: values.PAN, Data: values.Data, CVV: values.CVV, NAME: values.NAME,Status: values.Status}

	bankcardtype := services.BankCardVerification(values.PAN)
	if bankcardtype != "BRT card" {
		fmt.Println("some")
		return
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

func pay(ctx *gin.Context){
    
	var errResp models.ErrorStruct
	//var req string
	

	//APIURL := "http://"+"localhost"+":8080"+"/preCheckwithNum"
	client := http.Client{Timeout: settings.AppSettings.AppParams.TimeoutReq2NPCSec * time.Second}
   // url:="http://localhost:6060/info"
	url1:="http://ya.ru"
	var resp, err = client.Post(url1, "application/xml", nil)
	if err!=nil{
		errResp.Status=404
		errResp.StatusDetails=err.Error()
		ctx.JSON(http.StatusBadRequest,errResp)
		return
	}

	// var response models.PaymentResponse

	var resString, _ = ioutil.ReadAll(resp.Body)
    req:=services.RemoveNonXMLCharacters(string(resString))
	fmt.Println(req)

	response,err:=services.JsonType()
	if err!=nil{
		errResp.Status=404
		errResp.StatusDetails=err.Error()
		ctx.JSON(http.StatusBadRequest,errResp)
		return
	}
	ctx.JSON(http.StatusOK,response)
	// form := url.Values{
	// 	"agentLogin":        {agentLogin},
	// 	"agentPassword":     {agentPassword},
	// 	"cardOperationType": {cardOperationType}}

	// rsp,err:=http.Get(APIURL)
	// if err!=nil{
	// 	panic(err)
	// }
	// defer rsp.Body.Close()
	// fmt.Println(rsp.Body)
	// ctx.String(200,"",rsp.Body)

	// APIURL := settings.AppSettings.HSMService.HSMUrl + info
	// agentLogin := ctx.PostForm("agentLogin")
	// agentPassw := ctx.PostForm("agentPassword")
	// cardHash := ctx.PostForm("cardHash")
	// clientCode := ctx.PostForm("clientCode")
	// reason := ctx.PostForm("reason")
	// agentTransID := ctx.PostForm("transID")

}
