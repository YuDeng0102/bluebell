package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-app/logger"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {

		c.String(http.StatusOK, "Hello World")
	})

	//r.POST("/register", controller.RegisterHanndler)

	v1 := r.Group("/api")
	{
		GetUserRoutes(v1)
	}
	return r
}
