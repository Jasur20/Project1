package db

import (
	"brt_adapter/models"
	"brt_adapter/settings"
	"context"
	"errors"
	//"fmt"
	"time"
)

func FindCardsBankWithNum(Phone string) (models.PreCheckInfo, int, error) {

	preCheckInfo := &models.PreCheckInfo{}
	db := InitPostgresDB()
	defer db.Close()

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	list := models.List{}

	rows := db.QueryRowContext(ctx, "select * from list where phone=$1", Phone)
	err := rows.Scan(&list.Phone, &list.PAN, &list.Data, &list.CVV, &list.NAME, &list.Status)

	if err != nil {
		
		return *preCheckInfo, 400, errors.New("Error with DB service!!!")
	}
	if list.Status != true {
		preCheckInfo.NAME=list.NAME
		return *preCheckInfo, 400, errors.New("cards blocked!!")
	}

	preCheckInfo.NAME = settings.FullNameToInitials(list.NAME)
	return *preCheckInfo, 200, nil

	// rows:=db.QueryRow("select *from list where phone=$1",Phone)
	// card:=models.List{}
	// err:=rows.Scan(&card.Phone, &card.PAN,&card.Data,&card.CVV,&card.NAME)
	// if err!=nil{
	// 	return nil,err
	// }
	// fmt.Println("hello")
	// fmt.Println(rows)
	// fmt.Println(card.Phone)
	// fmt.Println(card.NAME)
	// return rows,nil
}

func CheckID(tranID string)error{
	db:=InitPostgresDB()
	defer db.Close()

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	values := models.PaymentResponse{}
	rows:= db.QueryRowContext(ctx, "select *from paymentH where tranid=$1", tranID)
	err:= rows.Scan(&values.TransID,&values.Result,&values.ResultCode,&values.CardHolder,&values.CardNumberHash,&values.MaskedPanNumber,&values.Reason,&values.Amount,&values.PaymentURL,&values.Decline,&values.ApprovalCode)
	if err!=nil{
		return nil
	}
	return errors.New("repetition tranid!!")
}