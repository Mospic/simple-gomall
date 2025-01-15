package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(service map[string]interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.InitMiddleware(service))

	v1 := ginRouter.Group("/gomall")
	{
		//user
		user := v1.Group("/user")
		{
			user.POST("/register/", handlers.Register)
			user.POST("/login/", handlers.Login)
			user.GET("/", handlers.UserInfo)
		}
	}
	return ginRouter
}
