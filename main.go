package main

import (
	"apps-transaction/config"
	"apps-transaction/handler"
	"apps-transaction/repository"
	"apps-transaction/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectionDatabase()
	fmt.Println("API Running on port 8080")

	router := gin.Default()

	// Setup repository, service, and handler Product
	productRepo := repository.NewProductRepository(config.DB)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// api for product
	router.POST("/products", productHandler.CreateProduct)
	router.GET("/products/:id", productHandler.GetProductByID)
	router.GET("/products", productHandler.GetAllProducts)
	router.POST("/products/:id", productHandler.UpdateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)

	// port RUN!
	router.Run(":8080")
}
