package routes

import (
	"github.com/Jateq/ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.AddProductAdmin())
	incomingRoutes.GET("/users/productview", controllers.AllProducts())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
