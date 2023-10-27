package server

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	// enable CORS
	r.Use(CORSMidware())

	// register two type ping
	r.GET("/ping", Ping)
	r.POST("/ping", Ping)

	// configure cookie handler
	setupCookieRouter(r.Group(""))

	// just login
	r.POST("/login", Login)
	r.GET("/auth", Auth)

}
