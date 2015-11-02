package json

import "github.com/arschles/flexwork"

// ErrorJSON is an error to be encoded to JSON and returned. See constants in this package
// for some commonly used ones
type ErrorJSON struct {
	Code    int    `json:"-"`
	Type    string `json:"error"`
	Message string `json:"msg"`
}

const jsonEncodeErr = `{
  "error":"json",
  "msg":"couldn't encode JSON indicating the error"
}`

// Error returns a flexwork.Response that encodes a JSON error response
func Error(err ErrorJSON) flexwork.Response {
	return Encode(err.Code, err)
}
