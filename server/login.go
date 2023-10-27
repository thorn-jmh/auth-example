package server

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	param := struct {
		Username string `json:"username"`
	}{}

	if err := c.BindJSON(&param); err != nil {
		c.JSON(400, gin.H{
			"mgs": "bad request",
		})
		return
	}

	doLogin(c, param.Username)

	c.JSON(200, gin.H{
		"mgs": "success~",
	})
}

func Auth(c *gin.Context) {

	c.JSON(200, gin.H{
		"mgs": "hi~" + doAuth(c),
	})
}

// -----------------  impl --------------------

func doLogin(c *gin.Context, username string) {
}

func doAuth(c *gin.Context) string {
	return ""
}
