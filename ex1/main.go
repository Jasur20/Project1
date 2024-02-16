package main

import (
	//"brt_adapter/db"
	//"brt_adapter/models"
	"brt_adapter/routes"

	//"brt_adapter/routes"
	//"context"
	//"fmt"
	//"time"

	//"brt_adapter/models"
	//"brt_adapter/routes"
	"brt_adapter/settings"
	//"fmt"
	//"strings"
	//"regexp"
)

func main() {
	settings.AppSettings=settings.ReadSettings()
	routes.Init()
}
