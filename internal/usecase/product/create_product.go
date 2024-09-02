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
	ProductRepository product.ProductRepository
}

func (u *CreateProductUsecase) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error){
	product := product.NewProduct(input.Name, input.Price)
	err := u.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}
	return &CreateProductOutputDto{
			ID: product.ID,
			Name: product.Name,
			Price: product.Price,
	}, nil
}