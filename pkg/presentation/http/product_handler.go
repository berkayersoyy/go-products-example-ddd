package http

import (
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/tracing/jaeger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
	"net/http"
	"strconv"
)

//productHandler Product handler
type productHandler struct {
	ProductService domain.ProductService
}

//ProvideProductAPI Provide product api
func ProvideProductAPI(p domain.ProductService) domain.ProductHandler {
	return &productHandler{ProductService: p}
}

// @BasePath /api/v1

// GetAllProducts
// @Summary Fetch all product
// @Schemes
// @Description Fetch all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} domain.Product
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/ [get]
func (p *productHandler) GetAllProducts(c *gin.Context) {
	tracer, span := jaeger.InitJaeger(c, "ProductHandler.GetAllProducts", "GET")
	err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		span.Finish()
		log.Fatalf("Error %s", err)
	}
	products := p.ProductService.GetAllProducts(c)
	span.SetTag("http.status_code", "200")
	span.Finish()
	c.JSON(http.StatusOK, gin.H{"products": products})
}

// @BasePath /api/v1

// GetProductByID
// @Summary Fetch product by id
// @Schemes
// @Description Fetch product by id
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} domain.Product
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/{id} [get]
func (p *productHandler) GetProductByID(c *gin.Context) {
	tracer, span := jaeger.InitJaeger(c, "ProductHandler.GetProductByID", "GET")
	err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		span.Finish()
		log.Fatalf("Error %s", err)
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.GetProductByID(c, uint(id))
	if product == (domain.Product{}) {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "404")
		span.Finish()
		c.JSON(http.StatusNotFound, "Product not found")
		return
	}
	span.SetTag("http.status_code", "200")
	span.Finish()
	c.JSON(http.StatusOK, gin.H{"product": domain.ToProductDTO(product)})
}

// @BasePath /api/v1

// AddProduct
// @Summary Add Product
// @Schemes
// @Description Add Product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body domain.Product true "Product ID"
// @Success 200 {object} domain.Product
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/ [post]
func (p *productHandler) AddProduct(c *gin.Context) {
	tracer, span := jaeger.InitJaeger(c, "ProductHandler.AddProduct", "POST")
	err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		span.Finish()
		log.Fatalf("Error %s", err)
	}
	var product domain.Product
	err = c.BindJSON(&product)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		span.Finish()
		c.Status(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(product)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		span.Finish()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	createdProduct := p.ProductService.AddProduct(c, product)
	span.SetTag("http.status_code", "201")
	span.Finish()
	c.JSON(http.StatusCreated, gin.H{"product": domain.ToProductDTO(createdProduct)})
}

// @BasePath /api/v1

// UpdateProduct
// @Summary Update Product
// @Schemes
// @Description Update Product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body domain.ProductDTO true "Product ID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/ [put]
func (p *productHandler) UpdateProduct(c *gin.Context) {
	tracer, span := jaeger.InitJaeger(c, "ProductHandler.UpdateProduct", "UPDATE")
	err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		span.Finish()
		log.Fatalf("Error %s", err)
	}
	var productDTO domain.ProductDTO
	err = c.BindJSON(&productDTO)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		span.Finish()
		c.Status(http.StatusBadRequest)
		return
	}

	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		span.Finish()
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.GetProductByID(c, uint(id))
	if product == (domain.Product{}) {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		span.Finish()
		c.Status(http.StatusBadRequest)
		return
	}

	product.Name = productDTO.Name
	product.Price = productDTO.Price
	product.Description = productDTO.Description
	p.ProductService.AddProduct(c, product)

	span.SetTag("http.status_code", "201")
	span.Finish()
	c.Status(http.StatusCreated)
}

// @BasePath /api/v1

// DeleteProduct
// @Summary Delete Product
// @Schemes
// @Description Delete Product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/{id} [delete]
func (p *productHandler) DeleteProduct(c *gin.Context) {
	tracer, span := jaeger.InitJaeger(c, "ProductHandler.DeleteProduct", "DELETE")
	err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		span.Finish()
		log.Fatalf("Error %s", err)
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.GetProductByID(c, uint(id))
	if product == (domain.Product{}) {
		ext.LogError(span, err)
		span.SetTag("http.status_code", "400")
		span.Finish()
		c.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.DeleteProduct(c, product)

	span.SetTag("http.status_code", "201")
	span.Finish()
	c.Status(http.StatusCreated)
}
