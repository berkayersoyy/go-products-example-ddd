package application

import (
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
)

type userServiceDynamoDb struct {
	UserRepository domain.UserRepositoryDynamoDb
}

//ProvideUserServiceDynamoDb Provide user service via dynamodb
func ProvideUserServiceDynamoDb(u domain.UserRepositoryDynamoDb) domain.UserServiceDynamoDb {
	return userServiceDynamoDb{UserRepository: u}
}

func (u userServiceDynamoDb) Update(ctx *gin.Context, user domain.User) error {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
	span := tracer.StartSpan("UserServiceDynamoDb.Update", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	err = u.UserRepository.Update(ctx, user)
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		return err
	}
	span.Finish()
	return nil
}
func (u userServiceDynamoDb) FindByUUID(ctx *gin.Context, uuid string) (domain.User, error) {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
	span := tracer.StartSpan("UserServiceDynamoDb.FindByUUID", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	user, err := u.UserRepository.FindByUUID(ctx, uuid)
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		return domain.User{}, err
	}
	span.Finish()
	return user, nil
}
func (u userServiceDynamoDb) FindByUsername(ctx *gin.Context, username string) (domain.User, error) {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
	span := tracer.StartSpan("UserServiceDynamoDb.FindByUsername", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	user, err := u.UserRepository.FindByUsername(ctx, username)
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		return domain.User{}, err
	}
	span.Finish()
	return user, nil
}

func (u userServiceDynamoDb) Insert(ctx *gin.Context, user domain.User) error {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
	span := tracer.StartSpan("UserServiceDynamoDb.Insert", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	err = u.UserRepository.Insert(ctx, user)
	if err != nil {
		return err
	}
	span.Finish()
	return nil
}
func (u userServiceDynamoDb) Delete(ctx *gin.Context, id string) error {
	tracer := opentracing.GlobalTracer()
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
	span := tracer.StartSpan("UserServiceDynamoDb.Delete", ext.RPCServerOption(parentSpan))
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		log.Printf("Error %s", err)
	}
	err = u.UserRepository.Delete(ctx, id)
	if err != nil {
		ext.LogError(span, err)
		span.Finish()
		return err
	}
	span.Finish()
	return nil
}
func (u userServiceDynamoDb) CreateTable(ctx *gin.Context) error {
	err := u.UserRepository.CreateTable(ctx)
	if err != nil {
		return err
	}
	return nil
}
