package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"rest-mang/database"
	"rest-mang/helpers"
	"rest-mang/middlewares"
	"rest-mang/models"
	"rest-mang/routes"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())

	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}
