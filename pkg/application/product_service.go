package application

import "github.com/berkayersoyy/go-products-example-ddd/pkg/domain"

type productService struct {
	ProductRepository domain.ProductRepository
}

// ProvideProductService Provide product service
func ProvideProductService(p domain.ProductRepository) domain.ProductService {
	return &productService{ProductRepository: p}
}

func (p *productService) GetAllProducts() []domain.Product {
	return p.ProductRepository.GetAllProducts()
}

func (p *productService) GetProductByID(id uint) domain.Product {
	return p.ProductRepository.GetProductByID(id)
}

func (p *productService) AddProduct(product domain.Product) domain.Product {
	p.ProductRepository.AddProduct(product)

	return product
}

func (p *productService) DeleteProduct(product domain.Product) {
	p.ProductRepository.DeleteProduct(product)
}
