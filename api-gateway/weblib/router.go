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
			user.GET("/", handlers.UserInfo)
			user.POST("/register/", handlers.Register)
			user.POST("/login/", handlers.Login)
			user.POST("/update", handlers.Update)
			user.POST("/delete", handlers.DeleteUser)
			user.POST("/logout", handlers.Logout)
		}
		//token
		token := v1.Group("/token")
		{
			token.POST("/varify", handlers.VarifyToken)
		}

		// product
		product := v1.Group("/product")
		{
			product.GET("/list/", handlers.ListProducts)
			product.GET("/", handlers.ProductInfo)
			product.GET("/search/", handlers.SearchProducts)

			product.POST("/create/", handlers.CreateProducts)
			product.POST("/delete/", handlers.DeleteProducts)
			//product.POST("/update/", handlers.UpdateProducts)
		}
	}
	return ginRouter
}
