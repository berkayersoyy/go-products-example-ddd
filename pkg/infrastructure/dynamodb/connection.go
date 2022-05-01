package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/infrastructure/config"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/infrastructure/errors"
	awstrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/aws/aws-sdk-go/aws"
	"log"
	"time"
)

const requestCanceledError = "RequestCanceled"

type DynamoDB struct {
	Connection *dynamodb.DynamoDB
	Timeout    time.Duration
}
type DynamoConnection interface {
	RunQuery(ctx context.Context, input *dynamodb.QueryInput) (*dynamodb.QueryOutput, domain.Error)
	GetItem(ctx context.Context, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, domain.Error)
	BatchGetItem(ctx context.Context, input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, domain.Error)
	PutItem(ctx context.Context, input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, domain.Error)
	DeleteItem(ctx context.Context, input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, domain.Error)
}

func New(cfg config.Config) *DynamoDB {
	dynamo := &DynamoDB{
		Timeout:    cfg.Timeout,
		Connection: newConnection(cfg.Region, cfg.AccessKey, cfg.SecretKey, cfg.EndpointUrl),
	}

	return dynamo
}
func newConnection(region string, accessKey string, secretKey string, endpointURL string) *dynamodb.DynamoDB {
	awsConfig := aws.Config{
		Region:      &region,
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:    &endpointURL,
	}

	sess, err := session.NewSession(&awsConfig)
	if err != nil {
		log.Fatal(err)
	}

	return dynamodb.New(awstrace.WrapSession(sess))
}
func (d *DynamoDB) DeleteItem(ctx context.Context, input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, domain.Error) {
	ctx, cancelFn := d.createCtxWithTimeout(ctx)
	defer cancelFn()

	output, err := d.Connection.DeleteItemWithContext(ctx, input)
	if err != nil {
		return nil, handleAWSError(err)
	}

	return output, nil
}
func (d *DynamoDB) RunQuery(ctx context.Context, input *dynamodb.QueryInput) (*dynamodb.QueryOutput, domain.Error) {
	ctx, cancelFn := d.createCtxWithTimeout(ctx)
	defer cancelFn()

	output, err := d.Connection.QueryWithContext(ctx, input)
	if err != nil {
		return nil, handleAWSError(err)
	}

	return output, nil
}
func (d *DynamoDB) GetItem(ctx context.Context, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, domain.Error) {
	ctx, cancelFn := d.createCtxWithTimeout(ctx)
	defer cancelFn()

	output, err := d.Connection.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, handleAWSError(err)
	}

	return output, nil
}
func (d *DynamoDB) PutItem(ctx context.Context, input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, domain.Error) {
	ctx, cancelFn := d.createCtxWithTimeout(ctx)
	defer cancelFn()

	output, err := d.Connection.PutItemWithContext(ctx, input)
	if err != nil {
		return nil, handleAWSError(err)
	}

	return output, nil
}
func (d *DynamoDB) BatchGetItem(ctx context.Context, input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, domain.Error) {
	ctx, cancelFn := d.createCtxWithTimeout(ctx)
	defer cancelFn()

	output, err := d.Connection.BatchGetItemWithContext(ctx, input)
	if err != nil {
		return nil, handleAWSError(err)
	}

	return output, nil
}
func (d *DynamoDB) createCtxWithTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, d.Timeout)
}
func handleAWSError(err error) domain.Error {
	if awsErr, ok := err.(awserr.Error); ok {
		if awsErr.Code() == requestCanceledError {
			return errors.New(errors.DynamoDBTimeoutError, errors.DynamoDBTimeout)
		}

		return errors.New(awsErr.Code(), awsErr.Error())
	}

	return errors.New(errors.DataError, err.Error())
}
