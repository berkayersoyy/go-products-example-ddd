package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/joho/godotenv"
	"github.com/twinj/uuid"
	"log"
	"os"
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
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.DeletedAt = nil
	UUID := uuid.NewV4()
	user.UUID = UUID.String()
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

func (u userRepository) Find(ctx context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"uuid": {S: aws.String(id)},
		},
	}
	//TODO dynamodb ayarla ve docker compose ayarla docker hub icin ve api testler ayarlar
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

func (u userRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"uuid": {S: aws.String(id)},
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
			"uuid": {S: aws.String(user.UUID)},
		},
		ExpressionAttributeNames: map[string]*string{
			"#username":  aws.String("username"),
			"#password":  aws.String("password"),
			"#UpdatedAt": aws.String("UpdatedAt"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":username":  {S: aws.String(user.Username)},
			":password":  {S: aws.String(user.Password)},
			":UpdatedAt": {S: aws.String(time.Now().Format(time.RFC3339))},
		},
		UpdateExpression: aws.String("set #username = :username, #password = :password, #UpdatedAt = :UpdatedAt"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	if _, err := u.client.UpdateItemWithContext(ctx, input); err != nil {
		log.Println(err)

		return domain.ErrInternal
	}

	return nil
}

//New Returns new
func New() (*session.Session, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return session.NewSessionWithOptions(
		session.Options{
			Config: aws.Config{
				Credentials:      credentials.NewStaticCredentials(os.Getenv("DynamoDBID"), os.Getenv("DynamoDBSECRET"), ""),
				Region:           aws.String(os.Getenv("DynamoDBREGION")),
				S3ForcePathStyle: aws.Bool(true),
			},
			Profile: os.Getenv("DynamoDBPROFILE"),
		},
	)
}
