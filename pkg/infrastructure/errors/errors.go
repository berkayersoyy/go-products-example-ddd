package errors

import "github.com/berkayersoyy/go-products-example-ddd/pkg/domain"

const (
	DataError            = "data_error"
	BadRequest           = "bad_request"
	DynamoDBTimeoutError = "dynamodb_timeout_error"
	DynamoDBTimeout      = "DynamoDB timeout"
)

type appError struct {
	message   string
	errorType string
	data      interface{}
}

func New(errorType string, message string, data ...interface{}) domain.Error {
	return &appError{errorType: errorType, message: message, data: data}
}

func (e *appError) Error() string {
	return e.message
}

func (e *appError) ErrorType() string {
	return e.errorType
}

func (e *appError) Data() interface{} {
	return e.data
}
