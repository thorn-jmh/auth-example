package main

import (
    "github.com/gin-gonic/gin"
)

func SetupCookieRouter(r *gin.Engine) {
    cookie := r.Group("/cookie")
    cookie.Use(CORSMidware())
    {
        cookie.GET("/simple", CookieSimpleHandler)
        cookie.GET("/delete", DeleteCookieHandler)
        cookie.GET("/httponly", CookieWithHttpOnlyHandler)
        cookie.GET("/secure", CookieWithSecureHandler)
    }
}

func CookieSimpleHandler(c *gin.Context) {
    c.SetCookie("test", "qwq", 3600, "/", "localhost", false, false)
    c.String(200, "Cookie set")
}

func DeleteCookieHandler(c *gin.Context) {
    c.SetCookie("test", "qwq", -1, "/", "localhost", false, false)
    c.String(200, "Cookie deleted")
}

func CookieWithHttpOnlyHandler(c *gin.Context) {
    c.SetCookie("test", "qwq", 3600, "/", "localhost", false, true)
    c.String(200, "Cookie set")
}

func CookieWithSecureHandler(c *gin.Context) {
    c.SetCookie("test", "qwq", 3600, "/", "localhost", true, false)
    c.String(200, "Cookie set")
}
