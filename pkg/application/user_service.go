package application

import (
	"context"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
	"net/http"
)

type userService struct {
	UserRepository domain.UserRepository
}

//ProvideUserService Provide user service via dynamodb
func ProvideUserService(u domain.UserRepository) domain.UserService {
	return userService{UserRepository: u}
}

func (u userService) Update(ctx context.Context, user domain.User) error {
	tracer := opentracing.GlobalTracer()
	header := ctx.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("UserService.Update", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	err = u.UserRepository.Update(ctx, user)
	if err != nil {
		ext.LogError(span, err)
		return err
	}
	return nil
}
func (u userService) FindByUUID(ctx context.Context, uuid string) (domain.User, error) {
	tracer := opentracing.GlobalTracer()
	header := ctx.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("UserService.FindByUUID", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	user, err := u.UserRepository.FindByUUID(ctx, uuid)
	if err != nil {
		ext.LogError(span, err)
		return domain.User{}, err
	}
	return user, nil
}
func (u userService) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	tracer := opentracing.GlobalTracer()
	header := ctx.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("UserService.FindByUsername", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	user, err := u.UserRepository.FindByUsername(ctx, username)
	if err != nil {
		ext.LogError(span, err)
		return domain.User{}, err
	}
	return user, nil
}

func (u userService) Insert(ctx context.Context, user domain.User) error {
	tracer := opentracing.GlobalTracer()
	header := ctx.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("UserService.Insert", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	err = u.UserRepository.Insert(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func (u userService) Delete(ctx context.Context, id string) error {
	tracer := opentracing.GlobalTracer()
	header := ctx.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("UserService.Delete", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	err = u.UserRepository.Delete(ctx, id)
	if err != nil {
		ext.LogError(span, err)
		return err
	}
	return nil
}
func (u userService) CreateTable(ctx context.Context) error {
	err := u.UserRepository.CreateTable(ctx)
	if err != nil {
		return err
	}
	return nil
}
