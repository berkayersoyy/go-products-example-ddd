package domain

type Error interface {
	Error() string
	ErrorType() string
	Data() interface{}
}
