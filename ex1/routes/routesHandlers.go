package routes

import (
	"brt_adapter/db"
	"brt_adapter/models"
	"brt_adapter/services"
	//"errors"
	"fmt"

	//"strings"

	//"brt_adapter/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func preCheck(ctx *gin.Context){

	var errResp models.ErrorStruct
	var values models.BankCard

	if err:=ctx.ShouldBindJSON(&values); err!=nil{
		errResp.Status=400
		errResp.StatusDetails="Bad type input "
		ctx.JSON(http.StatusBadRequest,errResp)
		return
	}
    
	err:=services.Validation(values.PAN,values.Data,values.CVV,values.NAME)
	if err!=nil{
		errResp.Status=400
		errResp.StatusDetails=err.Error()
		ctx.JSON(http.StatusBadRequest,errResp)
		return 
	}
	
	values=models.BankCard{PAN: values.PAN,Data: values.Data,CVV: values.CVV, NAME: values.NAME}
    fmt.Println(values.CVV)
	
	bankcardtype:=services.BankCardVerification(values.PAN)
	if bankcardtype!="BRT card"{
		fmt.Println("some")
		return
	}
	fmt.Println("BRT card")
	rows,err:=db.FindCardsBank(values.PAN,values.Data,values.CVV,values.NAME)
	if err!=nil{
		errResp.Status=404
		errResp.StatusDetails=err.Error()
	}
	fmt.Println(rows)
	
	//response:=services.CheckCard()
}