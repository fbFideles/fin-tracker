package controllers

import "github.com/gin-gonic/gin"

func UserRegister(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "user created",
	})
}

func UserGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "top top top",
	})
}
