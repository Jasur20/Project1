package controller

import (
	"alif/internal/integration"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)
//Добавить общий парсинг respBody


type EndPointController interface {
	InitRoutes()
}

type endPointController struct {
	engine      *gin.Engine
	integration integration.Integration
	token       string
}

func (e endPointController) InitRoutes() {
	//e.engine.GET("/service", e.GetServiceList)
	e.engine.POST("/payment", e.Payment)
	e.engine.POST("/payment/status", e.PaymentStatus)
	e.engine.POST("/receiver/info", e.ReceiverInfo)
}


func NewEndpointController(token string, engine *gin.Engine, integration integration.Integration) EndPointController {
	return &endPointController{token: token, engine: engine, integration: integration}
}


func (e *endPointController) Payment(ctx *gin.Context){
	token:=ctx.GetHeader("X-Token")
	operID:=ctx.GetHeader("X-Request-ID")
	if token !=e.token{
		logrus.WithFields(logrus.Fields{
			"X-Token":      token,
			"X-Request-ID": operID,
		}).Warn("bad token")
		ctx.JSON(http.StatusBadRequest, "Bad token")
		return
	}

	var paymentReq integration.PaymentRequestBody
	body,err:=io.ReadAll(ctx.Request.Body)
	if err!=nil{
		logrus.New().WithError(err).Warn("on request reading")
		ctx.JSON(http.StatusBadRequest,"Bad request")
		return
	}
	logrus.WithFields(
		logrus.Fields{
			"X-token": token,
			"X-oper-id": operID,
			"body": string(body),
		},
	).Info("Request to endpoint")

	err=json.Unmarshal(body,&paymentReq)
	if err!=nil{
		logrus.WithError(err).Warn("on request reading")
		ctx.JSON(http.StatusBadRequest,"Bad body")
		return
	}
	resp:=e.integration.Payment(operID,&paymentReq)
	respBody,err:=json.Marshal(resp)
	if err!=nil{
		logrus.WithError(err).Warn("on marshaling response")
		ctx.JSON(http.StatusBadRequest,"Bad body")
		return 
	}
	logrus.WithFields(
		logrus.Fields{
			"body": string(respBody),
		},
	).Info("response from endpoint")
	ctx.JSON(http.StatusOK,resp)
}

func (e *endPointController) ReceiverInfo(ctx *gin.Context) {
	token := ctx.GetHeader("X-token")
	operID := ctx.GetHeader("X-oper-id")
	fmt.Println("token", token)
	fmt.Println("e.token", e.token)
	if token != e.token {
		ctx.JSON(http.StatusBadRequest, "Bad token")
		return
	}
	var req integration.GetReceiverInfoRequestBody
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logrus.WithError(err).Warn("on request reading")
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	logrus.WithFields(
		logrus.Fields{
			"X-token":   token,
			"X-oper-id": operID,
			"body":      string(body),
		},
	).Info("Request to endpoint")

	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println("error: ", err)
		ctx.JSON(http.StatusBadRequest, "Bad body")
		return
	}

	resp:=e.integration.ReceiverInfo(operID,&req)
	respBody,err:=json.Marshal(resp)
	if err!=nil{
		logrus.WithError(err).Warn("on marshaling response")
		ctx.JSON(http.StatusBadRequest,"Bad body")
		return
	}
	logrus.WithFields(
		logrus.Fields{
			"body":string(respBody),
		},
	)
	ctx.JSON(http.StatusOK,resp)
}

func (e *endPointController) PaymentStatus(c *gin.Context) {
	token := c.GetHeader("X-token")
	operID := c.GetHeader("X-oper-id")

	if token != e.token {
		c.JSON(http.StatusBadRequest, "Bad token")
		return
	}
	var req integration.PostCheckRequestBody

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.WithError(err).Warn("on request reading")
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	logrus.WithFields(
		logrus.Fields{
			"X-token":   token,
			"X-oper-id": operID,
			"body":      string(body),
		},
	).Info("Request to endpoint")
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println("error: ", err)
		c.JSON(http.StatusBadRequest, "Bad body")
		return
	}

	resp := e.integration.PostCheck(operID, &req)
	respBody, err := json.Marshal(resp)
	if err != nil {
		logrus.WithError(err).Warn("on marshaling response")
		c.JSON(http.StatusBadRequest, "Bad body")
		return
	}
	logrus.WithFields(
		logrus.Fields{
			"body": string(respBody),
		},
	).Info("response from endpoint")
	c.JSON(http.StatusOK, resp)
}