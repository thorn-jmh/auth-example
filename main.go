package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"login/server"
)

func main() {
	readConfig()
	StartServer()
}

// -----------------  helper --------------------

func readConfig() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err)
	}
}

func StartServer() {
	r := gin.Default()
	server.InitRouter(r)
	logrus.Fatal(r.Run(":" + viper.GetString("ServePort")))
}
