package routes

import (
	"github.com/gin-gonic/gin"
	"web-app/controller"
)

func GetPostRoutes(router *gin.RouterGroup) {

	userGroup := router.Group("/post")
	{

		userGroup.GET("/:id", controller.GetPostHandler)
		userGroup.GET("", controller.GetPostList2Handler)
		//userGroup.GET("2", controller.GetPostList2Handler)
		userGroup.Use(controller.JWTAuthMiddleware())
		userGroup.POST("/vote", controller.VoteHandler)
		userGroup.POST("/", controller.CreatePostHandler)

	}
}
