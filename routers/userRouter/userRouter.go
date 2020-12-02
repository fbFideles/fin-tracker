package userRouter

import (
	"github.com/fbFideles/fin-tracker/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("", controllers.UserGetAll)
	r.POST("register", controllers.UserRegister)
}
