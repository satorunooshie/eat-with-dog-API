package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

type AppError interface {
	Error() string
	Unwrap() error
	Format(s fmt.State, v rune)
	FormatError(p xerrors.Printer) error
	Add(field string, data interface{}) AppError

	BadRequest() AppError
	Unauthorized() AppError
	NotFound() AppError
	InternalServerError() AppError

	Panic() AppError
	Critical() AppError
	Warn() AppError
	Info() AppError
	IsPanic() bool
	IsCritical() bool
	IsWarn() bool
	IsInfo() bool

	New(s ...string) AppError
	Errorf(format string, args ...interface{}) AppError
	Wrap(err error, s ...string) AppError
	Wrapf(err error, format string, args ...interface{}) AppError
	Code() string
	Status() int
	InfoMessage() string
	IsServerError() bool
	Messagef(args ...interface{}) AppError
	Is(err error) bool
}
