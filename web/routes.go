package web

import (
	"github.com/gin-gonic/gin"
)

func initRoutes(ginEngine *gin.Engine) {
	ginEngine.GET("/go", start)

	reportRoutes := ginEngine.Group("/go/api")
	reportRoutes.GET("/", ok)
}