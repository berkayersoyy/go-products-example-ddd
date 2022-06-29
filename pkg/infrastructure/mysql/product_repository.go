package mysql

import (
	"context"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
	"net/http"
)

type productRepository struct {
	DB *gorm.DB
}

//ProvideProductRepository Provide product repository
func ProvideProductRepository(DB *gorm.DB) domain.ProductRepository {
	return &productRepository{DB: DB}
}

func (p *productRepository) GetAllProducts(c context.Context) []domain.Product {
	tracer := opentracing.GlobalTracer()
	header := c.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("ProductRepository.GetAllProducts", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	var products []domain.Product
	p.DB.Find(&products)

	return products
}

func (p *productRepository) GetProductByID(c context.Context, id uint) domain.Product {
	tracer := opentracing.GlobalTracer()
	header := c.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("ProductRepository.GetProductByID", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	var product domain.Product
	p.DB.First(&product, id)

	return product
}

func (p *productRepository) AddProduct(c context.Context, product domain.Product) domain.Product {
	tracer := opentracing.GlobalTracer()
	header := c.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("ProductRepository.AddProduct", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	p.DB.Save(&product)

	return product
}

func (p *productRepository) DeleteProduct(c context.Context, product domain.Product) {
	tracer := opentracing.GlobalTracer()
	header := c.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("ProductRepository.GetAllProducts", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	p.DB.Delete(&product)
}
