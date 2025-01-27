package routes

import (
	"github.com/gin-gonic/gin"
	"web-app/controller"
)

func GetUserRoutes(router *gin.RouterGroup) {
	router.POST("/register", controller.RegisterHanndler) // 用户注册
	router.POST("/login", controller.LoginHanndler)       // 用户登陆
	// 													````
	userGroup := router.Group("/user", controller.JWTAuthMiddleware())
	{
		userGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
	}
}
