package errors

import (
	"fmt"
	"net/http"

	"golang.org/x/xerrors"

	"github.com/satorunooshie/eat-with-dog-API/server/library/i18n"
)

func (e *appError) New(s ...string) AppError {
	m := e.code
	if len(s) != 0 {
		m = s[0]
	}
	return e.new(m)
}

func (e *appError) Errorf(format string, args ...interface{}) AppError {
	return e.New(fmt.Sprintf(format, args...))
}

func (e *appError) Wrap(err error, s ...string) AppError {
	m := e.Code()
	if len(s) != 0 {
		m = s[0]
	}
	ne := e.new(m)
	ne.next = err
	return ne
}

func (e *appError) Wrapf(err error, format string, args ...interface{}) AppError {
	ne := e.new(fmt.Sprintf(format, args...))
	ne.next = err
	return ne
}

func (e *appError) Messagef(args ...interface{}) AppError {
	e.infoMessage = fmt.Sprintf(e.infoMessage, args...)
	return e
}

func (e *appError) Code() string {
	if e.code != `` {
		return e.code
	}
	if next := AsAppError(e.next); next != nil {
		return next.Code()
	}
	return `not_defined`
}

func (e *appError) Status() int {
	if e.status != 0 {
		return e.status
	}
	if next := AsAppError(e.next); next != nil {
		return next.Status()
	}
	return http.StatusInternalServerError
}

func (e *appError) InfoMessage() string {
	if e.infoMessage != `` {
		return e.InfoMessage()
	}
	if next := AsAppError(e.next); next != nil {
		return next.InfoMessage()
	}
	return defaultErrorMessage
}

func (e *appError) IsServerError() bool {
	return e.Status() >= http.StatusInternalServerError
}

func (e *appError) new(s string) *appError {
	e.message = s
	e.frame = xerrors.Caller(2)
	return e
}

func newError(code, s string) *appError {
	return &appError{
		code:        code,
		infoMessage: s,
	}
}

func newBadRequest(code, s string) AppError {
	e := newError(code, i18n.T(s))
	e.status = http.StatusBadRequest
	_ = e.Info()
	return e
}

func newUnauthorized(code, s string) AppError {
	e := newError(code, i18n.T(s))
	e.status = http.StatusUnauthorized
	_ = e.Info()
	return e
}

func newForbidden(code, s string) AppError {
	e := newError(code, i18n.T(s))
	e.status = http.StatusForbidden
	_ = e.Info()
	return e
}

func newConflict(code, s string) AppError {
	e := newError(code, i18n.T(s))
	e.status = http.StatusConflict
	_ = e.Info()
	return e
}

func newNotFound(code, s string) AppError {
	e := newError(code, i18n.T(s))
	e.status = http.StatusNotFound
	_ = e.Info()
	return e
}

func newInternalServerError(code, s string) AppError {
	e := newError(code, i18n.T(s))
	e.status = http.StatusInternalServerError
	_ = e.Critical()
	return e
}
