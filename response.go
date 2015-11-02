package flexwork

import "net/http"

// Response is the interface that is returned by a ReturnHandler's ServeHTTP func.
// it is intended to hold all data necessary to respond to an HTTP request
type Response interface {
	Finish(w http.ResponseWriter)
}

type funcResp struct {
	handler func(http.ResponseWriter)
	headers http.Header
}

// ResponseFunc is a convenience wrapper for the Response interface
type ResponseFunc func(w http.ResponseWriter)

func (r ResponseFunc) Finish(w http.ResponseWriter) {
	r(w)
}

// Error returns a Response that wraps http.Error
func Error(code int, msg string) Response {
	return ResponseFunc(func(w http.ResponseWriter) {
		http.Error(w, msg, code)
	})
}

// Redirect returns a Response that wraps http.Redirect
func Redirect(code int, location string, r *http.Request) Response {
	return ResponseFunc(func(w http.ResponseWriter) {
		http.Redirect(w, r, location, code)
	})
}

// Code returns a Response that wraps w.WriteHeader
func Code(code int) Response {
	return ResponseFunc(func(w http.ResponseWriter) {
		w.WriteHeader(code)
	})
}
