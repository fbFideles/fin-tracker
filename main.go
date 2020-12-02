package main

import (
	"github.com/fbFideles/fin-tracker/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		routers.Router(v1)
	}

	r.Run(":3000")
}
