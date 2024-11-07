package repository

import (
	"database/sql"
	"github.com/jalvess021/api-golang/api/internal/domain/product"
)

type ProductRepositoryPgsql struct {
	DB *sql.DB
}

func NewProductRepositoryPgsql(db *sql.DB) *ProductRepositoryPgsql {
	return &ProductRepositoryPgsql{DB: db}
}

func (r *ProductRepositoryPgsql) Create(product *product.ProductEntity) error {
	query := "INSERT INTO products (id, name, price) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, product.ID, product.Name, product.Price)

	if err != nil{
		return err
	}

	return nil
}

func (r *ProductRepositoryPgsql) FindAll() ([]*product.ProductEntity, error) {
	rows, err := r.DB.Query("select * from products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*product.ProductEntity
	for rows.Next() {
		var product product.ProductEntity
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}