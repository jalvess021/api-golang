package usecase

import "github.com/jalvess021/api-golang/api/internal/domain/product"

type ListProductsOutputDto struct{
	ID string
	Name string
	Price float64
}

type ListProductsUseCase struct{
	ProductRepositoryInterface product.ProductRepositoryInterface
}

func NewListProductsUseCase(productRepositoryInterface product.ProductRepositoryInterface) *ListProductsUseCase {
	return &ListProductsUseCase{ProductRepositoryInterface: productRepositoryInterface}
}

func (u *ListProductsUseCase) Execute() ([]*ListProductsOutputDto, error) {
	products, err := u.ProductRepositoryInterface.FindAll()
	if err != nil {
		return nil, err
	}

	var productsOutput []*ListProductsOutputDto
	for _, product := range products{
		productsOutput = append(productsOutput, &ListProductsOutputDto{
			ID: product.ID,
			Name: product.Name,
			Price: product.Price,
		})
	}
	return productsOutput, nil
}