package utils;

import (
	"fmt"
)
const (
	DbError = iota
	InsuficientLimitError = iota
)

type QueryError struct {
	Code int
	Msg string
	Err error
}

func (qe QueryError) Error() string{
	return fmt.Sprintf("%s\nErrorCode:%d\nMsg:%s", qe.Err.Error(), qe.Code, qe.Msg)
}
