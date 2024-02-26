package routes

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() {

	lumberLog := &lumberjack.Logger{
		Filename:   "",/*settings.AppSettings.AppParams.LogFile*/
		MaxSize:    10, // megabytes
		MaxBackups: 100,
		MaxAge:     100, //days
	}

	gin.DefaultWriter=io.MultiWriter(os.Stdout,lumberLog)
	log.SetOutput(gin.DefaultWriter)

	routes:=gin.Default()
    routes.POST("/preCheckwithNum",preCheckwithNum)
	routes.POST("/preCheckwithPan",preCheckwithPan)
	routes.POST("/Payment",pay)
	routes.GET("/postCheck",postCheck)

	routes.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound,gin.H{"error":"NOT_FOUND"})
	})

	routes.Run(":8080")
}