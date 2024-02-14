package main

import (
	//"brt_adapter/db"
	"brt_adapter/routes"
	"brt_adapter/settings"
	//"strings"
	//"regexp"
)

func main() {
	settings.AppSettings=settings.ReadSettings()
	// err:=db.InitOracleDB()

	// if err!=nil{
	// 	log.Fatalln(err)
	// }
	routes.Init()


}
