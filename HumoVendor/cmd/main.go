package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
	"vendor/internal/adapter"
	"vendor/internal/controller"
	"vendor/internal/integration"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Appsetings struct {
	AppParams Params  `json:"app"`
	HumoPay   HumoPay `json:"humoPay"`
}

type Params struct {
	ServerName string `json:"serverName"`
	PortRun    int    `json:"portRun"`
	LogFile    string `json:"logFile"`
	ServerURL  string `json:"serverURL"`
	Token      string `json:"token"`
	TimeOut    int `json:"timeOut"`
}

type HumoPay struct {
	Address    string `json:"address"`
	Secret     string `json:"secret"`
	PsId       string `json:"ps_id"`
	BarkiTojik []int  `json:"barki_tojik"`
	Vodokanal  []int  `json:"vodokanal"`
}

func ReadConfig() Appsetings {
	viper.SetConfigName("settings.json")
	viper.SetConfigType("json")
	viper.AddConfigPath("C:/Users/safarzoda_j/HumoVendor")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var config Appsetings
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return config
}

func customClient(timeout time.Duration) *http.Client {

	transport := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   1 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	client := http.Client{
		Transport: &transport,
		Timeout:   timeout * time.Second,
	}

	return &client
}

func main() {

	var config Appsetings
	config = ReadConfig()
	fmt.Println(config)
	httpClient := customClient(time.Duration(config.AppParams.TimeOut))
	adapter := adapter.NewAdapter(config.HumoPay.Address, config.HumoPay.PsId, httpClient, config.HumoPay.Secret)
	integration := integration.NewIntegration(adapter)

	engine := gin.Default()
	endpoint := controller.NewEndpointController(config.AppParams.Token, engine, integration)
	endpoint.InitRoutes()
	engine.Run(config.AppParams.ServerURL + ":" + fmt.Sprint(config.AppParams.PortRun))
}
