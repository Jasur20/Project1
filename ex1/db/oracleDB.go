package db

import (
	//"brt_adapter/settings"
	"database/sql"
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