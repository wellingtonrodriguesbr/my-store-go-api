package repository

import (
	"database/sql"
	"fmt"
	"my-store-api-go/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, price FROM products"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int

	query, err := pr.connection.Prepare("INSERT INTO products" + "(name, price)" + " VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM products WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &product, nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	query, err := pr.connection.Prepare("DELETE FROM products WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return err
	}

	var product model.Product

	err = query.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			return err
		}

		return err
	}

	query.Close()

	return nil
}

func (pr *ProductRepository) UpdateProduct(id int, product model.Product) (*model.Product, error) {
	query, err := pr.connection.Prepare("UPDATE products SET name = $2, price = $3 WHERE id = $1 RETURNING id, name, price")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	updatedProduct := model.Product{}

	err = query.QueryRow(id, product.Name, product.Price).Scan(&updatedProduct.ID, &updatedProduct.Name, &updatedProduct.Price)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &updatedProduct, nil
}
