package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/soyornaldo/ecomerce/controllers"
	"github.com/soyornaldo/ecomerce/database"
	"github.com/soyornaldo/ecomerce/middleware"
	"github.com/soyornaldo/ecomerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = 8000
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
