package main

import (
	"alif/internal/adapter"
	"alif/internal/controller"
	"alif/internal/integration"

	"fmt"
	"net"
	"net/http"
	"time"

	//"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


type Config struct{
	EndPoint EndPoint `mapstructure:"endpoint"`
	Adapter Adapter `mapstructure:"adapter"`
}

type EndPoint struct{
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Token string `mapstructure:"token"`
}

type Adapter struct{
	Alif Alif `mapstructure:"alif"`
}

type Alif struct{
	Url string `mapstructure:"url"`
	UserID string `mapstructure:"userid"`
	Secret string `mapstructure:"secret"`
	TimeOut int `mapstructure:"timeout"`
}


var Settings Config

func ReadConfig() Config{
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return config
}



func customClient(timeout time.Duration) *http.Client {
	//ref: Copy and modify defaults from https://golang.org/src/net/http/transport.go
	//Note: Clients and Transports should only be created once and reused
	transport := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			// Modify the time to wait for a connection to establish
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
	var config Config
	config=ReadConfig()
	httpClient:=customClient(time.Duration(config.Adapter.Alif.TimeOut))
	adapter:=adapter.NewAdapter(config.Adapter.Alif.Url,config.Adapter.Alif.UserID,httpClient)
	integration:=integration.NewIntegration(adapter)
	
	engine:=gin.Default()
	endpoint:=controller.NewEndpointController(config.EndPoint.Token,engine,integration)
	endpoint.InitRoutes()
	engine.Run(config.EndPoint.Host+":"+fmt.Sprint(config.EndPoint.Port))
}