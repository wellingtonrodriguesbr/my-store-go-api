package controller

import (
	"database/sql"
	"my-store-api-go/model"
	"my-store-api-go/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product) 

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		response := model.Response{
			Message: "Product id not found",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "Failed to convert product id",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(productId) 

	if product == nil {
		response := model.Response{
			Message: "Product not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}


	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		response := model.Response{
			Message: "Product id not found",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "Failed to convert product id",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productUseCase.DeleteProduct(productId) 

	if err != nil {
		if err == sql.ErrNoRows {
			response := model.Response {
				Message: "Product not found",
			}
		
		ctx.JSON(http.StatusNotFound, response)
		return
	}
}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}