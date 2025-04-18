package web

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)

	ginEngine := gin.Default()

	initRoutes(ginEngine)

	ginEngine.Run(port)
}