package json

import "errors"

type Err error

func NewErr(str string) Err {
	return Err(errors.New(str))
}

var (
	ErrFormatting   = NewErr("json_formatting")
	ErrDB           = NewErr("database")
	ErrUnauthorized = NewErr("unauthorized")
	ErrBadRequest   = NewErr("bad request")
	ErrInternal     = NewErr("internal_error")
	ErrNotFound     = NewErr("not_found")
)

const (
	jsonEncodeErr = `{
    "error":"json",
    "msg":"couldn't encode JSON indicating the error"
  }`
)

type ErrorJSON struct {
	Type    string `json:"error"`
	Message string `json:"msg"`
}

func Error(code int, eType Err, msg string) flexwork.Response {
	err := ErrorJSON{Type: eType.Error(), Message: msg}
	return JSON(code, err)
}
