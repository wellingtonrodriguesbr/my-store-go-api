package usecase

import (
	"my-store-api-go/model"
	"my-store-api-go/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}
	
	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) GetProductById(productId int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(productId)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUseCase) DeleteProduct(productId int) (error) {
	err := pu.repository.DeleteProduct(productId)

	if err != nil {
		return err
	}

	return nil
}