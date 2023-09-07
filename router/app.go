package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "sends/docs"
	"sends/middlewares"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	//配置swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
