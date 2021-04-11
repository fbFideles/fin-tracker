package userHandler

import (
	"net/http"

	"github.com/fbFideles/fin-tracker/app/userApplication"
	"github.com/fbFideles/fin-tracker/models/userModel"

	"github.com/gin-gonic/gin"
)

func SingUp(c *gin.Context) {
	var req userModel.ReqUserSingUp
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := userApplication.SingUp(c, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, token)
}

func SingIn(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
	})
}
