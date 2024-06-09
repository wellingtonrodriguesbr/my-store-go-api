package routes

import (
	"database/sql"
	"my-store-api-go/controller"
	"my-store-api-go/repository"
	"my-store-api-go/usecase"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, dbConnection *sql.DB) *gin.RouterGroup {
	productRepo := repository.NewProductRepository(dbConnection)
	productUseCase := usecase.NewProductUseCase(productRepo)
	productController := controller.NewProductController(productUseCase)

	productRoutes := router.Group("/")
	{
		productRoutes.GET("/products", productController.GetProducts)
		productRoutes.GET("/products/:productId", productController.GetProductById)
		productRoutes.POST("/products", productController.CreateProduct)
		productRoutes.PATCH("/products/:productId", productController.UpdateProduct)
		productRoutes.DELETE("/products/:productId", productController.DeleteProduct)
	}

	return productRoutes
}
