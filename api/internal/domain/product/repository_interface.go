package product

type ProductRepositoryInterface interface{
	Create(product *ProductEntity) error
	FindAll() ([]*ProductEntity, error)
}