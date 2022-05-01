package domain

import "github.com/jinzhu/gorm"
import _ "github.com/go-playground/validator/v10"

// swagger:model Product
type Product struct {
	gorm.Model
	Name        string  `json:"name" validate:"required,min=2,max=45"`
	Price       float32 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

// swagger:model ProductDTO
type ProductDTO struct {
	ID          uint    `json:"id,string,omitempty"`
	Name        string  `json:"name" validate:"required,min=2,max=45"`
	Price       float32 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

func ToProduct(productDTO ProductDTO) Product {

	return Product{Name: productDTO.Name, Price: productDTO.Price, Description: productDTO.Description}
}

func ToProductDTO(product Product) ProductDTO {
	return ProductDTO{ID: product.ID, Price: product.Price, Name: product.Name, Description: product.Description}
}

func ToProductDTOs(products []Product) []ProductDTO {
	productdtos := make([]ProductDTO, len(products))

	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}

	return productdtos
}