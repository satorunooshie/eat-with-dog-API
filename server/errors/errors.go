package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

type appError struct {
	// 標準を満たすエラー
	next    error
	message string
	frame   xerrors.Frame

	data  []map[string]interface{}
	level level

	code        string
	infoMessage string
	status      int
}

func (e *appError) Error() string {
	next := AsAppError(e.next)
	if next != nil {
		return next.Error()
	}
	if e.next == nil {
		if e.message == `` {
			return `no message`
		}
		return e.message
	}
	return e.next.Error()
}

func (e *appError) Is(err error) bool {
	if er := AsAppError(err); er != nil {
		return e.Code() == er.Code()
	}
	return false
}

func (e *appError) Unwrap() error {
	return e.next
}

func (e *appError) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

func (e *appError) FormatError(p xerrors.Printer) error {
	var s string
	if e.level != "" {
		s += fmt.Sprintf("[%s] ", e.level)
	}
	if e.code != "" {
		s += fmt.Sprintf("[%s] ", e.code)
	}
	if e.message != "" {
		s += e.message
	}
	if len(e.data) != 0 {
		if s != "" {
			s += "\n"
		}
		s += fmt.Sprintf("data: %+v", e.data)
	}
	p.Print(s)
	e.frame.Format(p)
	return e.next
}

func (e *appError) Add(field string, data interface{}) AppError {
	if e.data == nil {
		e.data = make([]map[string]interface{}, 0)
	}
	e.data = append(e.data, map[string]interface{}{field: data})
	return e
}

func New(s string) AppError {
	return create(s)
}

func Errorf(format string, args ...interface{}) AppError {
	return create(fmt.Sprintf(format, args...))
}

func Wrap(err error, s ...string) AppError {
	if err == nil {
		return nil
	}

	var m string
	if len(s) != 0 {
		m = s[0]
	}
	e := create(m)
	e.next = err
	return e
}

func Wrapf(format string, err error, args ...interface{}) AppError {
	e := create(fmt.Sprintf(format, args...))
	e.next = err
	return e
}

func AsAppError(err error) *appError {
	if err == nil {
		return nil
	}
	var e *appError
	if As(err, &e) {
		return e
	}
	return nil
}

func As(err error, target interface{}) bool {
	return xerrors.As(err, target)
}

func create(s string) *appError {
	var e appError
	e.message = s
	e.frame = xerrors.Caller(2)
	return &e
}
