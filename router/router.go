package router

import (
	"photo-app/controllers"
	"photo-app/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.GET("/userstatus", controllers.GetLastUserStatus)
	}

	productRouter := r.Group("/product")
	{
		productRouter.Use(middlewares.Authentication())
		
		productRouter.GET("/", controllers.GetAllProduct)
		productRouter.POST("/", middlewares.ProductAuthorization(), controllers.CreateProduct)
		productRouter.PUT("/:photoId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:photoId", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	orderRouter := r.Group("/orders")
	{
		orderRouter.PATCH("/:orderid", controllers.UpdateItemOrder)
	}

	return r
}
