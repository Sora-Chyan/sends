package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "sends/docs"
	"sends/middlewares"
	"sends/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	//配置swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	admin := r.Group("/admin")
	admin.POST("/login", service.AdminLogin)
	change := r.Group("/admin")
	change.POST("/change", middlewares.AuthAdminCheck(), service.ChangeState)
	show := r.Group("/admin")
	show.GET("/show", middlewares.AuthAdminCheck(), service.ShowAll)
	return r
}
