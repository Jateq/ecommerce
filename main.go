package main

import (
	d "github.com/Jateq/ecommerce/database"
	"github.com/Jateq/ecommerce/middleware"
	"github.com/Jateq/ecommerce/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := d.NewApplication(d.ProductData(d.Client, "Products"), d.UserData(d.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/addtocart", app.AddToCart())
	router.GET("/remove", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}

//part 3
