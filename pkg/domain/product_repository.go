package domain

type ProductRepository interface {
	GetAllProducts() []Product
	GetProductByID(id uint) Product
	AddProduct(product Product) Product
	DeleteProduct(product Product)
}
