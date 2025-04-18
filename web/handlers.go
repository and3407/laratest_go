package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func start(ginContext *gin.Context) {
	fmt.Println("Start page")

	ginContext.JSON(http.StatusOK, "Start page")
}

func ok(ginContext *gin.Context) {
	fmt.Println("Api ok. Console Success.")

	ginContext.JSON(http.StatusOK, "Api OK. Json Success.")
}