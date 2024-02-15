package db

import (
	"brt_adapter/models"
	"database/sql"
	"errors"
	"fmt"
	//"fmt"
	//"fmt"
)


func FindCardsBank(PAN, Data, CVV, NAME string) (*sql.Row,error) {
	db := InitPostgresDB()
	defer db.Close()

	rows:= db.QueryRow("select * from card where pan=$1", PAN)
	if err != nil {
		return nil,errors.New("The card with the specified PAN was not found!!!")
	}
	card:=models.BankCard{}
	err:=rows.Scan(&card.PAN,&card.Data,&card.CVV,&card.NAME)
	if err!=nil{
		return nil,err
	}
	fmt.Println(card.PAN)
	return rows,nil

}