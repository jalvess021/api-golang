package usecase

import "github.com/jalvess021/api-golang/internal/domain/product"

type CreateProductInputDto struct{
	Name string
	Price float64
}

type CreateProductOutputDto struct{
	ID string
	Name string
	Price float64
}

type CreateProductUsecase struct {
	ProductRepositoryInterface product.ProductRepositoryInterface 
}

func NewCreateProductUseCase(ProductRepositoryInterface product.ProductRepositoryInterface) *CreateProductUsecase {
	return &CreateProductUsecase{ProductRepositoryInterface: ProductRepositoryInterface}
}

func (u *CreateProductUsecase) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error){
	product := product.NewProduct(input.Name, input.Price)
	err := u.ProductRepositoryInterface.Create(product)
	if err != nil {
		return nil, err
	}
	return &CreateProductOutputDto{
			ID: product.ID,
			Name: product.Name,
			Price: product.Price,
	}, nil
}