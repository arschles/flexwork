package flexwork

import "net/http"

// Response is the interface that is returned by a ReturnHandler's ServeHTTP func.
// it is intended to hold all data necessary to respond to an HTTP request
type Response interface {
	Finish(w http.ResponseWriter)
	Headers() http.Header
}

type funcResp struct {
	handler func(http.ResponseWriter)
	headers http.Header
}

// ResponseFunc is a convenience wrapper for a func representation of a Resonse.
// It's the analog for net/http's HandlerFunc
func ResponseFunc(f func(w http.ResponseWriter)) Response {
	return &funcResp{
		handler: f,
		headers: http.Header(map[string][]string{}),
	}
}

// Headers is the interface implementation
func (r *funcResp) Headers() http.Header {
	return r.headers
}

// Finish is the interface implementation
func (r *funcResp) Finish(w http.ResponseWriter) {
	// TODO: more efficient way to do this
	for name, vals := range r.headers {
		for _, val := range vals {
			w.Header().Add(name, val)
		}
	}
	r.handler(w)
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
