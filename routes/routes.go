package routes

import (
	"ecomerce/controllers"
	"github/gin-gonic/gin"
)

func userRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProducto())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
