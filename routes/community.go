package routes

import (
	"github.com/gin-gonic/gin"
	"web-app/controller"
)

func GetCommunityRoutes(router *gin.RouterGroup) {

	userGroup := router.Group("/community")
	{
		userGroup.GET("/category", controller.CommunityHandler)
		userGroup.GET("/:id", controller.CommunityDetailHandler)
	}
}
