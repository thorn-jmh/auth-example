package main

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
)

func main() {
    readConfig()
    StartServer()
}

func readConfig() {
    viper.SetConfigFile("config.yaml")
    err := viper.ReadInConfig()
    if err != nil {
        logrus.Fatal(err)
    }
}

func StartServer() {
    r := gin.Default()
    configureServer(r)

    logrus.Fatal(r.Run(":7777"))
}

func configureServer(r *gin.Engine) {

    SetupCookieRouter(r)
    SetupCORSRouter(r)
}
