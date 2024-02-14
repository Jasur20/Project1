package db

import (
	"brt_adapter/settings"
	"database/sql"
	"fmt"
)

var (
	err error
	ODB *sql.DB
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

func InitPostgresDB() error{
	postgresDBParams:=
}