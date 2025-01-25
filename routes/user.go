package routes

import (
	"github.com/gin-gonic/gin"
	"web-app/controller"
)

func GetUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", controller.RegisterHanndler) // 用户注册
		userGroup.POST("/login", controller.LoginHanndler)       // 用户登陆
	}
}
