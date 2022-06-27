package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Product Entity product
// swagger:model Product
type Product struct {
	gorm.Model
	Name        string  `json:"name" validate:"required,min=2,max=45"`
	Price       float32 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

//ProductDTO Product dto
// swagger:model ProductDTO
type ProductDTO struct {
	ID          uint    `json:"id,string,omitempty"`
	Name        string  `json:"name" validate:"required,min=2,max=45"`
	Price       float32 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

//ToProduct To_Product
func ToProduct(productDTO ProductDTO) Product {

	return Product{Name: productDTO.Name, Price: productDTO.Price, Description: productDTO.Description}
}

//ToProductDTO To_Product_Dto
func ToProductDTO(product Product) ProductDTO {
	return ProductDTO{ID: product.ID, Price: product.Price, Name: product.Name, Description: product.Description}
}

//ToProductDTOs To_Product_Dtos
func ToProductDTOs(products []Product) []ProductDTO {
	productdtos := make([]ProductDTO, len(products))

	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}

	return productdtos
}

//ProductHandler Product_Handler
type ProductHandler interface {
	GetAllProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
	AddProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

//ProductRepository Product_Repository
type ProductRepository interface {
	GetAllProducts(c *gin.Context) []Product
	GetProductByID(c *gin.Context, id uint) Product
	AddProduct(c *gin.Context, product Product) Product
	DeleteProduct(c *gin.Context, product Product)
}

//ProductService Product_Service
type ProductService interface {
	GetAllProducts(c *gin.Context) []Product
	GetProductByID(c *gin.Context, id uint) Product
	AddProduct(c *gin.Context, product Product) Product
	DeleteProduct(c *gin.Context, product Product)
}
