package main

import (
	"my-store-api-go/controller"
	"my-store-api-go/db"
	"my-store-api-go/repository"
	"my-store-api-go/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	productRepo := repository.NewProductRepository(dbConnection)
	productUseCase := usecase.NewProductUseCase(productRepo)
	productController := controller.NewProductController(productUseCase)

	server.GET("/products", productController.GetProducts)
	server.GET("/products/:productId", productController.GetProductById)
	server.POST("/products", productController.CreateProduct)
	server.PATCH("/products/:productId", productController.UpdateProduct)
	server.DELETE("/products/:productId", productController.DeleteProduct)

	server.Run(":8000")
}