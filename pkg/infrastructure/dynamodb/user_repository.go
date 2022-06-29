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
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/twinj/uuid"
	"log"
	"net/http"
	"os"
	"time"
)

type userRepository struct {
	Timeout time.Duration
	client  *dynamodb.DynamoDB
}

//ProvideUserRepository Provide user repository
func ProvideUserRepository(session *session.Session, Timeout time.Duration) domain.UserRepository {
	return userRepository{Timeout: Timeout, client: dynamodb.New(session)}
}

func (u userRepository) CreateTable(ctx context.Context) error {
	result, err := u.listTables(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	if contains(result.TableNames, "Users") {
		return nil
	}
	input := &dynamodb.CreateTableInput{
		TableName: aws.String("Users"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("UUID"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("UUID"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}

	out, err := u.client.CreateTableWithContext(ctx, input)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Successfully created table %s", out)
	return nil
}
func (u userRepository) Insert(ctx context.Context, user domain.User) error {
	tracer := opentracing.GlobalTracer()
	header := ctx.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("UserRepository.Insert", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.DeletedAt = nil
	UUID := uuid.NewV4()
	user.UUID = UUID.String()
	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		ext.LogError(span, err)
		log.Println(err)
		return domain.ErrInternal
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Users"),
		Item:      item,
		ExpressionAttributeNames: map[string]*string{
			"#UUID": aws.String("UUID"),
		},
		ConditionExpression: aws.String("attribute_not_exists(#UUID)"),
	}

	if _, err := u.client.PutItemWithContext(c, input); err != nil {
		log.Println(err)

		if _, ok := err.(*dynamodb.ConditionalCheckFailedException); ok {
			ext.LogError(span, err)
			return domain.ErrConflict
		}
		ext.LogError(span, err)
		return domain.ErrInternal
	}
	return nil
}

func (u userRepository) FindByUUID(ctx context.Context, id string) (domain.User, error) {
	tracer := opentracing.GlobalTracer()
	header := ctx.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("UserRepository.FindByUUID", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {S: aws.String(id)},
		},
	}

	res, err := u.client.GetItemWithContext(c, input)
	if err != nil {
		ext.LogError(span, err)
		log.Println(err)

		return domain.User{}, domain.ErrInternal
	}

	if res.Item == nil {
		ext.LogError(span, err)
		return domain.User{}, domain.ErrNotFound
	}

	var user domain.User
	if err := dynamodbattribute.UnmarshalMap(res.Item, &user); err != nil {
		ext.LogError(span, err)
		log.Println(err)
		return domain.User{}, domain.ErrInternal
	}
	return user, nil
}

func (u userRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	tracer := opentracing.GlobalTracer()
	header := ctx.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("UserRepository.FindByUsername", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.ScanInput{
		TableName:        aws.String("Users"),
		FilterExpression: aws.String("contains(Username, :Username)"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Username": {S: aws.String(username)},
		},
	}

	res, err := u.client.ScanWithContext(c, input)
	if err != nil {
		ext.LogError(span, err)
		log.Println(err)
		return domain.User{}, domain.ErrInternal
	}

	if *res.Count == 0 {
		ext.LogError(span, err)
		return domain.User{}, domain.ErrNotFound
	}

	var user domain.User
	for _, userItem := range res.Items {
		var userToScan domain.User
		err := dynamodbattribute.UnmarshalMap(userItem, &userToScan)
		if err != nil {
			log.Println(err)
			ext.LogError(span, err)
			return domain.User{}, domain.ErrInternal
		}
		if userToScan.Username != username {
			continue
		}
		user = userToScan
	}
	return user, nil
}

func (u userRepository) Delete(ctx context.Context, uuid string) error {
	tracer := opentracing.GlobalTracer()
	header := ctx.Value("header").(http.Header)
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("UserRepository.Delete", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {S: aws.String(uuid)},
		},
	}

	if _, err := u.client.DeleteItemWithContext(c, input); err != nil {
		ext.LogError(span, err)
		log.Println(err)

		return domain.ErrInternal
	}
	return nil
}

func (u userRepository) Update(ctx context.Context, user domain.User) error {
	tracer := opentracing.GlobalTracer()
	header := ctx.Value("header").(http.Header)
	//opentracing.StartSpanFromContext(ctx,"UserRepository.Update")
	parentSpan, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	span := tracer.StartSpan("UserRepository.Update", ext.RPCServerOption(parentSpan))
	defer span.Finish()
	if err != nil {
		ext.LogError(span, err)
		log.Printf("Error %s", err)
	}
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {S: aws.String(user.UUID)},
		},
		ExpressionAttributeNames: map[string]*string{
			"#Username":  aws.String("Username"),
			"#Password":  aws.String("Password"),
			"#UpdatedAt": aws.String("UpdatedAt"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Username":  {S: aws.String(user.Username)},
			":Password":  {S: aws.String(user.Password)},
			":UpdatedAt": {S: aws.String(time.Now().Format(time.RFC3339))},
		},
		UpdateExpression: aws.String("set #Username = :Username, #Password = :Password, #UpdatedAt = :UpdatedAt"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	if _, err := u.client.UpdateItemWithContext(c, input); err != nil {
		ext.LogError(span, err)
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
				Endpoint:         aws.String(os.Getenv("DynamoDBENDPOINTURL")),
			},
			Profile: os.Getenv("DynamoDBPROFILE"),
		},
	)
}
func contains(list []*string, compareItem string) bool {
	for _, listItem := range list {
		if *listItem == compareItem {
			return true
		}
	}
	return false
}
func (u userRepository) listTables(ctx context.Context) (*dynamodb.ListTablesOutput, error) {
	c, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	input := &dynamodb.ListTablesInput{}
	result, err := u.client.ListTablesWithContext(c, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
