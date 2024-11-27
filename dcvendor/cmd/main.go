package main

import (
	"dc_adapter/internal/adapter"
	"dc_adapter/internal/controller"
	"dc_adapter/internal/usecase"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppSettings struct {
	AppParams Params `yaml:"params"`
	DcPay   DcPay  `yaml:"dcpay"`
	Logs      Logs   `yaml:"logs"`
}

type Params struct {
	PortRun   int    `yaml:"portrun"`
	ServerURL string `yaml:"serverurl"`
	Token     string `yaml:"token"`
	TimeOut   int    `yaml:"timeout"`
}

type DcPay struct {
	Address  string `yaml:"address"`
	Login    string `yaml:"login"`
	Password string `yml:"password"`
}

type Logs struct {
	Dir  string `yaml:"dir"`
	Name string `yaml:"name"`
}

func ReadConfig() AppSettings {
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var config AppSettings
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

	config := ReadConfig()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	file, err := os.OpenFile(fmt.Sprintf("%s%s", config.Logs.Dir, config.Logs.Name), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Info("failed to set log output to file, using default stderr:" + err.Error())
	}
	logrus.Info("config: ", config)

	httpClient := customClient(time.Duration(config.AppParams.TimeOut))

	adapter := adapter.NewAdapter(config.DcPay.Address, config.DcPay.Login,httpClient, config.DcPay.Password)
	integration := usecase.NewIntegration(adapter)

	engine := gin.Default()
	endpoint := controller.NewEndpointController(config.AppParams.Token, engine, integration)
	endpoint.InitRoutes()
	engine.Run(config.AppParams.ServerURL + ":" + fmt.Sprint(config.AppParams.PortRun))

}
