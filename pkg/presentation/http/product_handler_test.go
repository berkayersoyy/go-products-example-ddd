package http

//
//import (
//	"bytes"
//	"encoding/json"
//	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
//	"github.com/berkayersoyy/go-products-example-ddd/pkg/mocks"
//	"github.com/gin-gonic/gin"
//	"github.com/jinzhu/gorm"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestProductHandler_GetAllProducts(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	t.Run("Success", func(t *testing.T) {
//		products := []domain.Product{{Name: "test-name", Price: 10, Description: "test-desc", Model: gorm.Model{ID: 1}}}
//		mockApi := mocks.ProductHandler{}
//		mockApi.On("GetAllProducts", mock.Anything).Return(products)
//
//		w := httptest.NewRecorder()
//		r := gin.Default()
//		r.GET("/v1/products", mockApi.GetAllProducts)
//		req, err := http.NewRequest(http.MethodGet, "/v1/products", nil)
//		req.Header.Set("Content-Type", "application/json")
//		assert.NoError(t, err)
//		r.ServeHTTP(w, req)
//
//		var respBody []domain.Product
//		err = json.NewDecoder(w.Body).Decode(&respBody)
//		assert.NoError(t, err)
//
//		assert.Equal(t, 200, w.Code)
//		assert.Equal(t, products, respBody)
//		mockApi.AssertNumberOfCalls(t, "GetAllProducts", 1)
//	})
//
//}
//func TestProductHandler_GetProductByID(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	t.Run("Success", func(t *testing.T) {
//		product := domain.Product{Name: "test-name", Price: 10, Description: "test-desc", Model: gorm.Model{ID: 1}}
//		mockApi := mocks.ProductHandler{}
//		mockApi.On("GetProductByID", mock.Anything).Return(product)
//
//		w := httptest.NewRecorder()
//		r := gin.Default()
//		r.Use(func(context *gin.Context) {
//			context.Set("id", uint(1))
//		})
//		r.GET("/v1/products/:id", mockApi.GetProductByID)
//		req, err := http.NewRequest(http.MethodGet, "/v1/products/1", nil)
//		req.Header.Set("Content-Type", "application/json")
//		assert.NoError(t, err)
//		r.ServeHTTP(w, req)
//
//		var respBody domain.Product
//		err = json.NewDecoder(w.Body).Decode(&respBody)
//		assert.NoError(t, err)
//
//		assert.Equal(t, 200, w.Code)
//		assert.Equal(t, product, respBody)
//		mockApi.AssertNumberOfCalls(t, "GetProductByID", 1)
//	})
//
//}
//func TestProductHandler_AddProduct(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	t.Run("Success", func(t *testing.T) {
//		product := domain.Product{Name: "test-name", Price: 10, Description: "test-desc", Model: gorm.Model{ID: 1}}
//		mockApi := mocks.ProductHandler{}
//		mockApi.On("AddProduct", mock.Anything).Return(product)
//
//		w := httptest.NewRecorder()
//		r := gin.Default()
//
//		reqBody, err := json.Marshal(product)
//		assert.NoError(t, err)
//		r.POST("/v1/product", mockApi.AddProduct)
//		req, err := http.NewRequest(http.MethodPost, "/v1/product", bytes.NewBuffer(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		assert.NoError(t, err)
//		r.ServeHTTP(w, req)
//
//		var respBody domain.Product
//		err = json.NewDecoder(w.Body).Decode(&respBody)
//		assert.NoError(t, err)
//
//		assert.Equal(t, 200, w.Code)
//		assert.Equal(t, product, respBody)
//		mockApi.AssertNumberOfCalls(t, "AddProduct", 1)
//	})
//}
//func TestProductHandler_UpdateProduct(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	t.Run("Success", func(t *testing.T) {
//		product := domain.Product{Name: "test-name", Price: 10, Description: "test-desc", Model: gorm.Model{ID: 1}}
//		mockApi := mocks.ProductHandler{}
//		mockApi.On("UpdateProduct", mock.Anything).Return(product)
//
//		w := httptest.NewRecorder()
//		r := gin.Default()
//
//		reqBody, err := json.Marshal(product)
//		assert.NoError(t, err)
//		r.PUT("/v1/products", mockApi.UpdateProduct)
//		req, err := http.NewRequest(http.MethodPut, "/v1/products", bytes.NewBuffer(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		assert.NoError(t, err)
//		r.ServeHTTP(w, req)
//
//		var respBody domain.Product
//		err = json.NewDecoder(w.Body).Decode(&respBody)
//		assert.NoError(t, err)
//
//		assert.Equal(t, 200, w.Code)
//		assert.Equal(t, product, respBody)
//		mockApi.AssertNumberOfCalls(t, "UpdateProduct", 1)
//	})
//}
//func TestProductHandler_DeleteProduct(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	t.Run("Success", func(t *testing.T) {
//		mockApi := mocks.ProductHandler{}
//		mockApi.On("DeleteProduct", mock.Anything)
//
//		w := httptest.NewRecorder()
//		r := gin.Default()
//		r.Use(func(context *gin.Context) {
//			context.Set("id", uint(1))
//		})
//
//		r.DELETE("/v1/products/:id", mockApi.DeleteProduct)
//		req, err := http.NewRequest(http.MethodDelete, "/v1/products/1", nil)
//		req.Header.Set("Content-Type", "application/json")
//		assert.NoError(t, err)
//		r.ServeHTTP(w, req)
//
//		assert.Equal(t, 200, w.Code)
//		mockApi.AssertNumberOfCalls(t, "DeleteProduct", 1)
//	})
//}
