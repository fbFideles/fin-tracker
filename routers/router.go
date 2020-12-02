package routers

import (
	"github.com/fbFideles/fin-tracker/routers/userRouter"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	userRouter.UserRouter(r.Group("user"))
}
