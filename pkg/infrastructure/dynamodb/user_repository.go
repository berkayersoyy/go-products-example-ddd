package dynamodb

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"log"
)

type userRepository struct {
	db DynamoConnection
}

func ProvideUserRepository(db DynamoConnection) domain.UserRepositoryCtx {
	return &userRepository{db: db}
}
func (u *userRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (u *userRepository) GetUserByID(ctx context.Context, id uint) (domain.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(string(id))},
		},
	}
	res, err := u.db.GetItem(ctx, input)
	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}
	if res.Item == nil {
		return domain.User{}, errors.New("Not Found")
	}
	var user domain.User
	er := dynamodbattribute.UnmarshalMap(res.Item, &user)
	if er != nil {
		log.Println(er)
		return domain.User{}, er
	}
	return user, nil
}
func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(username)},
		},
	}
	res, err := u.db.GetItem(ctx, input)
	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}
	if res.Item == nil {
		return domain.User{}, errors.New("Not Found")
	}
	var user domain.User
	er := dynamodbattribute.UnmarshalMap(res.Item, &user)
	if er != nil {
		log.Println(er)
		return domain.User{}, er
	}
	return user, nil
}

func (u *userRepository) AddUser(ctx context.Context, user domain.User) error {
	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Println(err)
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName:                aws.String("users"),
		Item:                     item,
		ExpressionAttributeNames: map[string]*string{"#id": aws.String("id")},
		ConditionExpression:      aws.String("attribute_not_exists(#id)"),
	}
	_, err = u.db.PutItem(ctx, input)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *userRepository) DeleteUser(ctx context.Context, user domain.User) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(string(user.ID))},
		},
	}
	_, err := u.db.DeleteItem(ctx, input)
	if err != nil {
		log.Println(err)
		return errors.New("Delete Failed")
	}
	return nil
}
