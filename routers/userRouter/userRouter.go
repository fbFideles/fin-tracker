package userRouter

import (
	"github.com/fbFideles/fin-tracker/handlers/userHandler"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	r.POST("singup", userHandler.SingUp)
}
