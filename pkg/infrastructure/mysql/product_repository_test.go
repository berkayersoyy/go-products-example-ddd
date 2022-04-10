package mysql

import (
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/mocks"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestProductRepository_GetAllProductsShouldReturnNotEmptyProductArray(t *testing.T) {
	products := []domain.Product{{Name: "test-product", Price: 10, Description: "test-desc", Model: gorm.Model{ID: 1}}}
	mockRepo := mocks.ProductRepository{}
	mockRepo.On("GetAllProducts").Return(products)

	resp := mockRepo.GetAllProducts()

	assert.Equal(t, products, resp)
	assert.NotEmpty(t, resp)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "GetAllProducts", 1)
}
func TestProductRepository_GetAllProductsShouldReturnEmptyProductArray(t *testing.T) {
	products := []domain.Product{}
	mockRepo := mocks.ProductRepository{}
	mockRepo.On("GetAllProducts").Return(products)

	resp := mockRepo.GetAllProducts()

	assert.Equal(t, products, resp)
	assert.Empty(t, resp)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "GetAllProducts", 1)
}
func TestProductRepository_GetProductByIDShouldReturnValidProduct(t *testing.T) {
	product := domain.Product{Name: "test-product", Price: 10, Description: "test-desc", Model: gorm.Model{ID: 1}}
	mockRepo := mocks.ProductRepository{}
	mockRepo.On("GetProductByID", mock.Anything).Return(product)

	resp := mockRepo.GetProductByID(uint(1))

	assert.Equal(t, product, resp)
	assert.NotEmpty(t, product, resp)
	assert.Equal(t, product.ID, resp.ID)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "GetProductByID", 1)
}
func TestProductRepository_GetProductByIDShouldReturnEmptyProduct(t *testing.T) {
	product := domain.Product{Name: "test-product", Price: 10, Description: "test-desc", Model: gorm.Model{ID: 1}}
	mockRepo := mocks.ProductRepository{}
	mockRepo.On("GetProductByID", mock.Anything).Return(domain.Product{})

	resp := mockRepo.GetProductByID(uint(0))

	assert.NotEqual(t, product, resp)
	assert.Empty(t, resp)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "GetProductByID", 1)
}
func TestProductRepository_AddProductShouldReturnValidProduct(t *testing.T) {
	product := domain.Product{Name: "test-product", Price: 10, Description: "test-desc", Model: gorm.Model{ID: 1}}
	mockRepo := mocks.ProductRepository{}
	mockRepo.On("AddProduct", mock.Anything).Return(product)

	resp := mockRepo.AddProduct(product)

	assert.NotEmpty(t, resp)
	assert.Equal(t, product, resp)
	assert.NotNil(t, resp)
	mockRepo.AssertNumberOfCalls(t, "AddProduct", 1)
}
func TestProductRepository_DeleteProduct(t *testing.T) {
	product := domain.Product{Name: "test-product", Price: 10, Description: "test-desc", Model: gorm.Model{ID: 1}}
	mockRepo := mocks.ProductRepository{}
	mockRepo.On("DeleteProduct", mock.Anything)

	mockRepo.DeleteProduct(product)

	mockRepo.AssertNumberOfCalls(t, "DeleteProduct", 1)
}
