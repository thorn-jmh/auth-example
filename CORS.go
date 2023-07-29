package main

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
)

func SetupCORSRouter(r *gin.Engine) {
    cors := r.Group("/cors")
    cors.Use(CORSMidware())
    {
        cors.GET("/ping", PingHandler)
        cors.POST("/ping", PostPingHandler)

        // Gin will not execute middleware and handlers if NOT FOUND
        // we need to add OPTIONS method to make it work.
        // In real world, we should use CORS config for all routes,
        // so we don't need to add OPTIONS method for each route.
        cors.OPTIONS("/ping", PingHandler)
    }
    r.GET("/ping", PingHandler)
    r.POST("/ping", PostPingHandler)
}

func PingHandler(c *gin.Context) {
    c.String(200, "pong")
}

func PostPingHandler(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}

func CORSMidware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Writer.Header().Set("Access-Control-Allow-Origin", getOrigin())
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(200)
        }
        c.Next()
    }
}

func getOrigin() string {
    origin := viper.GetString("RemoteOrigin")
    if origin == "" {
        origin = "*"
    }
    return origin
}
