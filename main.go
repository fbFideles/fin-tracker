package main

import (
	"github.com/fbFideles/fin-tracker/database"
	"github.com/fbFideles/fin-tracker/middleware"
	"github.com/fbFideles/fin-tracker/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	if err := database.NewConnection(); err != nil {
		panic(err)
	}
	defer database.CloseConnection()

	publicRouterGroup := r.Group("api/v1/public")
	{
		routers.PublicRouter(publicRouterGroup)
	}

	privateRouterGroup := r.Group("api/v1/private")
	privateRouterGroup.Use(
		middleware.Authorization(),
	)

	routers.PrivateRouter(privateRouterGroup)

	r.Run(":3000")
}
