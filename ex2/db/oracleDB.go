package db

import (
	//"brt_adapter/settings"
	//"brt_adapter/models"
	"brt_adapter/models"
	"database/sql"
	"errors"

	//"fmt"
	_"github.com/lib/pq"
	//"errors"
)

var (
	err error
	ODB *sql.DB
	PDB *sql.DB
)

// func InitOracleDB() error {
// 	oracleDBParams:= settings.AppSettings.OracleCFTDbParams
// 	oracleConnectionString := fmt.Sprintf("%s/%s@%s", oracleDBParams.User, oracleDBParams.Password, oracleDBParams.Server)

// 	ODB, err = sql.Open("godror", oracleConnectionString)
// 	if err!=nil{
// 		return err
// 	}
// 	return nil
// }
var (
	connStr= "host=localhost port=5432 user=postgres password=postgres dbname=test sslmode=disable"
)

func InitPostgresDB() *sql.DB{
	var err error
	PDB,err=sql.Open("postgres",connStr)
	if err!=nil{
		return nil
	}
	return PDB
}

func GetPostGresDB() *sql.DB{
	return PDB
}

func SavePay(tranid,result string,resultCode int,cardHolder,cardNumberHash,MaskedPanNumber,Reason string,Amount float64,PaymentURL,Decline,ApprovalCode string) error{
	db:=InitPostgresDB()
	defer db.Close()

	_,err:=db.Exec("insert into paymentH (tranid,result,resultcode,cardholder,cardnumberhash,maskedpannumber,reason,amount,paymenturl,decline,approvalcode) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)",
            tranid,result,resultCode,cardHolder,cardNumberHash,MaskedPanNumber,Reason,Amount,PaymentURL,Decline,ApprovalCode)
	if err!=nil{
		return errors.New("error with save process!!")
	}
	return nil
}

func PostCheck(tranid string) (*models.PostCheckResp,error){
	db:=InitPostgresDB()
	defer db.Close()
	resp := new(models.PostCheckResp)
	
	_,err:=db.Exec("select *from paymentH where tranid=$1",tranid)
	if err!=nil{
		resp.Status=400
		resp.StatusDetails=err.Error()
		return resp,errors.New("error with find tranid!!")
	}
	resp.Status=200
	resp.StatusDetails="Запрос успешно принят!!"
	return resp,nil
}
