package product

import (
	"github.com/google/uuid"
)

type ProductEntity struct {
	ID string
	Name string
	Price float64
}

func NewProduct(name string, price float64) *ProductEntity{
	return &ProductEntity{
		ID: uuid.New().String(),
		Name: name,
		Price: price,
	}
}