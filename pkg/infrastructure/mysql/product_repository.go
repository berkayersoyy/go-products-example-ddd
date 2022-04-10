package mysql

import (
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/jinzhu/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func ProvideProductRepository(DB *gorm.DB) domain.ProductRepository {
	return &productRepository{DB: DB}
}

func (p *productRepository) GetAllProducts() []domain.Product {
	var products []domain.Product
	p.DB.Find(&products)

	return products
}

func (p *productRepository) GetProductByID(id uint) domain.Product {
	var product domain.Product
	p.DB.First(&product, id)

	return product
}

func (p *productRepository) AddProduct(product domain.Product) domain.Product {
	p.DB.Save(&product)

	return product
}

func (p *productRepository) DeleteProduct(product domain.Product) {
	p.DB.Delete(&product)
}
