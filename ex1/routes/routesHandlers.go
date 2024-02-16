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
func preCheckwithNum(ctx *gin.Context){
	
	var list models.List
	var errResp models.ErrorStruct
	
	if err:=ctx.ShouldBindJSON(&list); err!=nil{
		errResp.Status=400
		errResp.StatusDetails="Bad type input"
		ctx.JSON(http.StatusBadRequest,errResp)
		return
	}
	
	
	err:=services.ValidationNum(list.Phone)
	if err!=nil{
		errResp.Status=400
		errResp.StatusDetails="error in verification Phone!!"
		ctx.JSON(http.StatusBadRequest,errResp)
		return
	}
    
	list=models.List{Phone: list.Phone}
	rows,status,err:=db.FindCardsBankWithNum(list.Phone)
	if err!=nil{
		fmt.Println("stop")
		errResp.Status=status
		errResp.StatusDetails=err.Error()
		ctx.JSON(http.StatusBadRequest,errResp)
		return
	}
	errResp.Status=status
	errResp.StatusDetails=err.Error()
	ctx.JSON(http.StatusOK,rows)	
}

func preCheckwithPan(ctx *gin.Context){

	var errResp models.ErrorStruct
	var values models.BankCard
	//var response models.PreCheckResp

	if err:=ctx.ShouldBindJSON(&values); err!=nil{
		errResp.Status=400
		errResp.StatusDetails="Bad type input "
		ctx.JSON(http.StatusBadRequest,errResp)
		return
	}

    
	err:=services.ValidationPan(values.PAN,values.Data,values.CVV,values.NAME)
	if err!=nil{
		errResp.Status=400
		errResp.StatusDetails=err.Error()
		ctx.JSON(http.StatusBadRequest,errResp)
		return 
	}
	
	values=models.BankCard{PAN: values.PAN,Data: values.Data,CVV: values.CVV, NAME: values.NAME}
	
	bankcardtype:=services.BankCardVerification(values.PAN)
	if bankcardtype!="BRT card"{
		fmt.Println("some")
		return
	}
	fmt.Println("BRT card")
	rows,status,err:=db.FindCardsBankWithPan(values.PAN)
	if err!=nil{
		errResp.Status=status
		errResp.StatusDetails=err.Error()
		ctx.JSON(http.StatusBadRequest,errResp)
		return
	}
	ctx.JSON(http.StatusOK,rows)

}