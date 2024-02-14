package routes

import (
	"brt_adapter/models"
	"brt_adapter/services"
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
	
	//response:=services.CheckCard()
}