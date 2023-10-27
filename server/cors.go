package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// -----------------  CORS midware --------------------

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
