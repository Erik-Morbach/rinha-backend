package utils

import "fmt"

type BaseError struct {
	Code int
	Msg  string
}

func (b BaseError) Error() string {
	return fmt.Sprint(b.Code) + " - " + b.Msg
}

const (
	DbError           = iota
	EmptyResultError  = iota
	RequestError      = iota
	ValidationError   = iota
	PaymentError      = iota
	UserNotFoundError = iota
)

// Just to simplify the error creation
func NewError(code int, msg string) error {
	return BaseError{Code: code, Msg: msg}
}

func VerifyErrorCode(err error) int {
	if v, ok := err.(BaseError); ok{
		return v.Code
	}
	return -1
}
