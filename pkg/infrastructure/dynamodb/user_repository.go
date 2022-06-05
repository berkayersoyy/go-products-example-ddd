package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/application/util/config"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"log"
	"strconv"
	"time"
)

type userRepository struct {
	Timeout time.Duration
	client  *dynamodb.DynamoDB
}

//ProvideUserRepository Provide user repository
func ProvideUserRepository(session *session.Session, Timeout time.Duration) domain.UserRepositoryCtx {
	return userRepository{Timeout: Timeout, client: dynamodb.New(session)}
}
func (u userRepository) Insert(ctx context.Context, user domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Println(err)
		return domain.ErrInternal
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("users"),
		Item:      item,
		ExpressionAttributeNames: map[string]*string{
			"#id": aws.String("id"),
		},
		ConditionExpression: aws.String("attribute_not_exists(#id)"),
	}

	if _, err := u.client.PutItemWithContext(ctx, input); err != nil {
		log.Println(err)

		if _, ok := err.(*dynamodb.ConditionalCheckFailedException); ok {
			return domain.ErrConflict
		}

		return domain.ErrInternal
	}

	return nil
}

func (u userRepository) Find(ctx context.Context, id uint) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(strconv.FormatUint(uint64(id), 10))},
		},
	}

	res, err := u.client.GetItemWithContext(ctx, input)
	if err != nil {
		log.Println(err)

		return domain.User{}, domain.ErrInternal
	}

	if res.Item == nil {
		return domain.User{}, domain.ErrNotFound
	}

	var user domain.User
	if err := dynamodbattribute.UnmarshalMap(res.Item, &user); err != nil {
		log.Println(err)

		return domain.User{}, domain.ErrInternal
	}

	return user, nil
}

func (u userRepository) Delete(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(strconv.FormatUint(uint64(id), 10))},
		},
	}

	if _, err := u.client.DeleteItemWithContext(ctx, input); err != nil {
		log.Println(err)

		return domain.ErrInternal
	}

	return nil
}

func (u userRepository) Update(ctx context.Context, user domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(strconv.FormatUint(uint64(user.ID), 10))},
		},
		ExpressionAttributeNames: map[string]*string{
			"#username": aws.String("username"),
			"#password": aws.String("password"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name":     {S: aws.String(user.Username)},
			":password": {N: aws.String(user.Password)},
		},
		UpdateExpression: aws.String("set #id = :id, #username = :username, #password = :password"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	if _, err := u.client.UpdateItemWithContext(ctx, input); err != nil {
		log.Println(err)

		return domain.ErrInternal
	}

	return nil
}
func New(config config.Config) (*session.Session, error) {
	return session.NewSessionWithOptions(
		session.Options{
			Config: aws.Config{
				Credentials:      credentials.NewStaticCredentials(config.ID, config.AccessSecret, ""),
				Region:           aws.String(config.Region),
				Endpoint:         aws.String(config.EndpointURL),
				S3ForcePathStyle: aws.Bool(true),
			},
			Profile: config.Profile,
		},
	)
}
