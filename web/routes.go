package web

import (
	"github.com/gin-gonic/gin"
)

func initRoutes(ginEngine *gin.Engine) {
	ginEngine.GET("/", start)

	reportRoutes := ginEngine.Group("/api")
	reportRoutes.GET("/", ok)
}