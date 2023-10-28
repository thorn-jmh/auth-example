package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

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
	user, err := doAuth(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"mgs": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"mgs": "hi~, " + user,
	})
}

// -----------------  impl --------------------

var users map[string]string = make(map[string]string)

func doLogin(c *gin.Context, username string) {
	uuid, _ := uuid.NewUUID()
	users[uuid.String()] = username
	c.SetCookie("token", uuid.String(), 3600, "/", "localhost", false, true)

}

func doAuth(c *gin.Context) (string, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return "", err
	} else if _, has := users[cookie]; has == false {
		return "", errors.New("no such user")
	} else {
		return users[cookie], nil
	}

}
