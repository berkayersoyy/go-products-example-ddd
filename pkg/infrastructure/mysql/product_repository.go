package mysql

import (
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
)

type productRepository struct {
	DB *gorm.DB
}

//ProvideProductRepository Provide product repository
func ProvideProductRepository(DB *gorm.DB) domain.ProductRepository {
	return &productRepository{DB: DB}
}

func (p *productRepository) GetAllProducts(c *gin.Context) []domain.Product {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span := tracer.StartSpan("ProductRepository.GetAllProducts", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	var products []domain.Product
	p.DB.Find(&products)
	span.Finish()

	return products
}

func (p *productRepository) GetProductByID(c *gin.Context, id uint) domain.Product {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span := tracer.StartSpan("ProductRepository.GetProductByID", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	var product domain.Product
	p.DB.First(&product, id)
	span.Finish()

	return product
}

func (p *productRepository) AddProduct(c *gin.Context, product domain.Product) domain.Product {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span := tracer.StartSpan("ProductRepository.AddProduct", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	p.DB.Save(&product)
	span.Finish()

	return product
}

func (p *productRepository) DeleteProduct(c *gin.Context, product domain.Product) {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span := tracer.StartSpan("ProductRepository.GetAllProducts", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	p.DB.Delete(&product)
	span.Finish()
}
