package services

import (
	"brt_adapter/models"
	"errors"
	//"log"
	"strconv"

	//"strconv"
	"strings"
	//"time"
)

func CheckCard(PAN, Data, CVV, NAME string) models.PreCheckResp{
	// var response models.PreCheckResp
	return models.PreCheckResp{}
}

//Проверка валидности даных карты
func Validation(PAN,Data,CVV,NAME string) (error){
    
	//Проверка PAN на корректность
	if len(strings.ReplaceAll(PAN," ",""))!=16{
		return errors.New("Error with PAN!!")
	}

	//Удаляем "/" из дата-строки
	dataStr:=strings.ReplaceAll(Data,"/","")

	//Выводим из даты-строки месяц
	mount,err:=strconv.Atoi(dataStr[:2])
	if err!=nil{
		return errors.New("Data should be digit!!")
	}

    //Выводим из даты-строки год
	year,err:=strconv.Atoi(dataStr[2:4])
	if err!=nil{
		return errors.New("Year should be digit")
	}

	//Проверка валидности даты-строки
	if len(dataStr)!=4{
		return errors.New("Error with Data!!")
	}
	if mount>12{
		return errors.New("The number of months does not exceed 12!!!")
	}
	if year>23{
		return errors.New("The Year must not exceed the present time!!")
	}
	
	//Проверка CVV на валидность
	if len(strings.ReplaceAll(CVV," ",""))!=3{
		return errors.New("Error with CVV!!")
	}
	if len(strings.ReplaceAll(NAME," ",""))==0{
		return errors.New("Error with NAME!!")
	}
	return nil

}