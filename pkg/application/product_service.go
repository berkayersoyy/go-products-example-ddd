package application

import (
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
)

type productService struct {
	ProductRepository domain.ProductRepository
}

// ProvideProductService Provide product service
func ProvideProductService(p domain.ProductRepository) domain.ProductService {
	return &productService{ProductRepository: p}
}

func (p *productService) GetAllProducts(c *gin.Context) []domain.Product {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span := tracer.StartSpan("ProductService.GetAllProducts", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	span.Finish()
	return p.ProductRepository.GetAllProducts(c)
}

func (p *productService) GetProductByID(c *gin.Context, id uint) domain.Product {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span := tracer.StartSpan("ProductService.GetProductByID", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	span.Finish()
	return p.ProductRepository.GetProductByID(c, id)
}

func (p *productService) AddProduct(c *gin.Context, product domain.Product) domain.Product {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span := tracer.StartSpan("ProductService.AddProduct", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	p.ProductRepository.AddProduct(c, product)
	span.Finish()
	return product
}

func (p *productService) DeleteProduct(c *gin.Context, product domain.Product) {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span := tracer.StartSpan("ProductService.DeleteProduct", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	p.ProductRepository.DeleteProduct(c, product)
	span.Finish()
}
