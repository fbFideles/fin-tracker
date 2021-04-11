package routers

import (
	"github.com/fbFideles/fin-tracker/routers/userRouter"
	"github.com/gin-gonic/gin"
)

// PublicRouter for routes that don't need
// authorization
func PublicRouter(r *gin.RouterGroup) {
	userRouter.UserRouter(r.Group("user"))
}

// PrivateRouter for routes that need
// authorization
func PrivateRouter(r *gin.RouterGroup) {}
