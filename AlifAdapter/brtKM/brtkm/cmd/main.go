package main

import (
	"brtkm/internal/adapter"
	"brtkm/internal/controller"
	"brtkm/internal/integration"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type EndPoint struct {
	Host  string `mapstructure:"host"`
	Port  int    `mapstructure:"port"`
	Token string `mapstructure:"token"`
}
type ISPC struct {
	URL          string `mapstructure:"url"`
	AgentLogin   string `mapstructure:"agent_login"`
	AgentPasword string `mapstructure:"agent_password"`
	Timeout      int    `mapstructure:"timeout"`
}
type CFT struct {
	URL     string `mapstructure:"url"`
	Timeout int    `mapstructure:"timeout"`
}
type Adapter struct {
	ISPC ISPC `mapstructure:"ispc"`
	CFT  CFT  `mapstructure:"cft"`
}

type Card struct {
	Pans []string `mapstructure:"pans"`
}

type Config struct {
	EndPoint EndPoint `mapstructure:"endpoint"`
	Adapter  Adapter  `mapstructure:"adapter"`
	Card     Card     `mapstructure:"card"`
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
	httpClient := customClient(time.Duration(config.Adapter.CFT.Timeout))
	adapter := adapter.NewAdapter(config.Adapter.CFT.URL, config.Adapter.ISPC.URL, config.Adapter.ISPC.AgentLogin, config.Adapter.ISPC.AgentPasword, config.Card.Pans, httpClient)
	integration := integration.NewIntegration(adapter)

	engine := gin.Default()
	endpoint := controller.NewEndpointController(config.EndPoint.Token, engine, integration)
	endpoint.InitRoutes()
	engine.Run(config.EndPoint.Host + ":" + fmt.Sprint(config.EndPoint.Port))
}
