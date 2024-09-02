package product

type ProductRepository interface{
	Create(product *ProductEntity) error
	FindAll() ([]*ProductEntity, error)
}