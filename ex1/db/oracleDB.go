package db

import (
	//"brt_adapter/settings"
	"database/sql"
	"errors"
	//"fmt"
	_ "github.com/lib/pq"
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

func SavePay(name,pan,tranid string,amount int,data string) error{
	db:=InitPostgresDB()
	defer db.Close()

	_,err:=db.Exec("insert into paymentH (name,pan,tranid,amount,data) values($1,$2,$3,$4,$5)",
	name,pan,tranid,amount,data)
	if err!=nil{
		return errors.New("error with save process!!")
	}
	return nil
}