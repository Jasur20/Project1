package db

import (
	//"brt_adapter/db"
	"brt_adapter/models"
	//"container/list"
	"database/sql"
	//"errors"
	"fmt"
	//"fmt"
	//"fmt"
)


func FindCardsBankWithPan(PAN, Data, CVV, NAME string) (*sql.Row,error) {
	db := InitPostgresDB()
	defer db.Close()

	rows:= db.QueryRow("select * from card where pan=$1", PAN)
	// if err != nil {
	// 	return nil,errors.New("The card with the specified PAN was not found!!!")
	// }
	card:=models.BankCard{}
	err:=rows.Scan(&card.PAN,&card.Data,&card.CVV,&card.NAME)
	if err!=nil{
		return nil,err
	}
	fmt.Println(card.PAN)
	return rows,nil

}

func FindCardsBankWithNum(Phone string) (*sql.Row,error){
	db:=InitPostgresDB()
	defer db.Close()

	rows:=db.QueryRow("select *from list where phone=$1",Phone)
	card:=models.List{}
	err:=rows.Scan(&card.Phone,&card.PAN,&card.Data,&card.CVV,&card.NAME)
	if err!=nil{
		return nil,err
	}
	fmt.Println(rows)
	fmt.Println(card.PAN)
	fmt.Println("not found")
	return rows,nil
}
