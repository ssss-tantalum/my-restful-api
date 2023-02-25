package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ssss.tantalum/my-restful-api/controllers"
	"github.com/ssss.tantalum/my-restful-api/middlewares"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}
