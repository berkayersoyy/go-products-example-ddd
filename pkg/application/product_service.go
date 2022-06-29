package application

import (
	"context"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
	"net/http"
)

type productService struct {
	ProductRepository domain.ProductRepository
}

// ProvideProductService Provide product service
func ProvideProductService(p domain.ProductRepository) domain.ProductService {
	return &productService{ProductRepository: p}
}

func (p *productService) GetAllProducts(c context.Context) []domain.Product {
	tracer := opentracing.GlobalTracer()
	header := c.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("ProductService.GetAllProducts", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	return p.ProductRepository.GetAllProducts(c)
}

func (p *productService) GetProductByID(c context.Context, id uint) domain.Product {
	tracer := opentracing.GlobalTracer()
	header := c.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("ProductService.GetProductByID", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	return p.ProductRepository.GetProductByID(c, id)
}

func (p *productService) AddProduct(c context.Context, product domain.Product) domain.Product {
	tracer := opentracing.GlobalTracer()
	header := c.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("ProductService.AddProduct", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	p.ProductRepository.AddProduct(c, product)
	return product
}

func (p *productService) DeleteProduct(c context.Context, product domain.Product) {
	tracer := opentracing.GlobalTracer()
	header := c.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("ProductService.DeleteProduct", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	p.ProductRepository.DeleteProduct(c, product)
}
