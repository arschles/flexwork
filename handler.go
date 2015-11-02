package flexwork

import "net/http"

// Handler is flexwork's analog of net/http's Handler. This version requires a Response to be returned.
// Returned responses can then be composed together, re-routed, duplicated, etc...
type Handler interface {
	ServeHTTP(r *http.Request) Response
}

// HandlerFunc is flexwork's analog of net/http's HandlerFunc
type HandlerFunc func(r *http.Request) Response

// ServeHTTP is the interface implementation
func (h HandlerFunc) ServeHTTP(r *http.Request) Response {
	return h(r)
}

// Wrapper is a net/http handler that wraps a ReturnHandler.
// Example usage:
//
//  http.Handle(flexwork.WrapFunc(func(r *http.Request) Response {
//    // do stuff
//    return flexwork.Code(http.StatusOK)
//  }))
type Wrapper struct {
	h Handler
}

// Wrap simply makes a Handler into a net/http Handler
func Wrap(h Handler) Wrapper {
	return Wrapper{h: h}
}

// WrapFunc simply makes a HandlerFunc into a net/http Handler
func WrapFunc(h HandlerFunc) Wrapper {
	return Wrapper{h: h}
}

// ServeHTTP is the interface implementation
func (a Wrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := a.h.ServeHTTP(r)
	res.Finish(w)
}
