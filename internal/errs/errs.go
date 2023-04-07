package errs

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidAccessToken     = errors.New("invalid auth token")
	ErrInvalidLoginOrPassword = errors.New("login or password is invalid")
	ErrLoginIsEmpty           = errors.New("login is empty")
	ErrLoginAlreadyExist      = errors.New("login already exist")
	ErrPasswordIsEmpty        = errors.New("password is empty")
)

type SQLError struct {
	Code string
	Err  error
}

func (se *SQLError) Error() string {
	return fmt.Sprintf("%v", se.Code)
}

func NewSQLError(code string) error {
	return &SQLError{
		Code: code,
	}
}

func (se *SQLError) Unwrap() error {
	return errors.New(se.Code)
}
