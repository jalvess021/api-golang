package repository

import (
	"database/sql"
	"github.com/jalvess021/api-golang/internal/domain/product"
)

type ProductRepositoryPgsql struct {
	DB *sql.DB
}

func NewProductRepositoryPgsql(db *sql.DB) *ProductRepositoryPgsql {
	return &ProductRepositoryPgsql{DB: db}
}

func (r *ProductRepositoryPgsql) Create(product *product.ProductEntity) error {
	_, err := r.DB.Exec("INSERT into products (id, name, price) values(?, ?, ?)", product.ID, product.Name, product.Price)

	if err != nil{
		return err
	}

	return nil
}

func (r *ProductRepositoryPgsql) findAll() ([]*product.ProductEntity, error) {
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