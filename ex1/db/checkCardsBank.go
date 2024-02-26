package db

import (
	//"brt_adapter/db"
	"brt_adapter/models"
	"brt_adapter/settings"
	"context"
	"errors"

	//"fmt"

	//"errors"
	"time"
)

func FindCardsBankWithPan(PAN string) (models.PreCheckInfo, int, error) {

	db := InitPostgresDB()
	defer db.Close()
	preCheckInfo := &models.PreCheckInfo{}
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	card := models.BankCard{}

	rows := db.QueryRowContext(ctx, "select *from card where pan=$1", PAN)
	// rows:=db.QueryRowContext(ctx,"select *from cards where pan=$1",PAN)
	err := rows.Scan(&card.PAN, &card.Data, &card.CVV, &card.NAME,&card.Status)
	if err != nil {
		return *preCheckInfo, 400, errors.New("error with db servise!!!")
	}
	if card.Status!=true{
		return *preCheckInfo,400,errors.New("cards blocked!!")
	}
	preCheckInfo.NAME = settings.FullNameToInitials(card.NAME)
	return *preCheckInfo, 200, nil
	// rows:= db.QueryRow("select * from card where pan=$1", PAN)
	// // if err != nil {
	// // 	return nil,errors.New("The card with the specified PAN was not found!!!")
	// // }
	// card:=&models.BankCard{}
	// err:=rows.Scan(&card.PAN,&card.Data,&card.CVV,&card.NAME)
	// if err!=nil{
	// 	return *card,403,errors.New("service is temporarily unavailable!!!")
	// }
	// fmt.Println(card.PAN)
	// return *card,200, nil

}

func FindCardsBankWithNum(Phone string) (models.PreCheckInfo, int, error) {

	preCheckInfo := &models.PreCheckInfo{}
	db := InitPostgresDB()
	defer db.Close()

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	list := models.List{}

	rows := db.QueryRowContext(ctx, "select * from list where phone=$1", Phone)
	err := rows.Scan(&list.Phone, &list.PAN, &list.Data, &list.CVV, &list.NAME,&list.Status)

	if err != nil {
		return *preCheckInfo, 400, errors.New("Error with DB service!!!")
	}
	if list.Status!=true{
		return *preCheckInfo,400,errors.New("cards blocked!!")
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
