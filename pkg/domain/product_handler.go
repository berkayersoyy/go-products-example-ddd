package domain

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	GetAllProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
	AddProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}
