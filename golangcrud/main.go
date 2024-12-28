package main

import (
	"golangcrud/controllers"
	"golangcrud/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	router := gin.Default()

	// Endpoint User
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Endpoint Produk
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products", controllers.GetProducts)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)

	// Endpoint Order
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)

	// Endpoint Payment
	router.POST("/payments", controllers.CreatePayment)

	router.Run(":8080")
}
