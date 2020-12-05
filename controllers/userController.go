package controllers

import (
	"github.com/fbFideles/fin-tracker/database"
	"github.com/fbFideles/fin-tracker/models"
	"github.com/gin-gonic/gin"
)

// UserRegister is a controllers wich defines the
// business logic for registering a user
func UserRegister(c *gin.Context) {
	msgErrPadrao := "Could not register user"
	user := models.User{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(400, map[string]interface{}{
			"message": msgErrPadrao + ": error on request read",
		})
		return
	}

	tx, err := database.NewTransaction(c)
	if err != nil {
		c.JSON(400, map[string]interface{}{
			"message": msgErrPadrao + ": error on opening database transaction",
		})
		return
	}
	defer tx.Rollback()

	if err = user.RegisterUser(tx); err != nil {
		c.JSON(400, map[string]interface{}{
			"message": msgErrPadrao + ": error on user database insert",
		})
		return
	}

	if err = tx.Commit(); err != nil {
		c.JSON(400, map[string]interface{}{
			"message": msgErrPadrao + ": error on user database insert commit",
		})
		return
	}

	c.JSON(201, user.ID)
}

// UserGetAll is a controller wich defines
// the business logic for getting all users
func UserGetAll(c *gin.Context) {
	c.JSON(400, map[string]interface{}{
		"message": "in progress...",
	})
}
