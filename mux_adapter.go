package flexwork

import "net/http"

// TraditionalMuxer is the interface that http.ServeMux & Gorilla Mux Router automatically adhere
// to. Drop one of those into the WrapMux func to create a flexwork-compatible API that
// uses that muxer under the covers
type Muxer interface {
	Handle(string, http.Handler)
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
}

// MuxAdapter can adapt any Muxer to be flexwork compatible, while still adhering
// to the net/http Handler interface. Sample usage:
//
//  adapter := flexwork.WrapMux(http.NewServeMux())
//  adapter.HandleFunc("/flexwork", func(r *http.Request) flexwork.Response {
//    // ...
//    return http.Code(http.StatusOK)
//  })
//  http.ListenAndServe(":8080", adapter)
type MuxAdapter struct {
	m Muxer
}

// WrapMux wraps m and returns the resulting MuxWrapper
func WrapMux(m Muxer) *MuxWrapper {
	return &MuxWrapper{m: m}
}

// Handle is the equivalent to http.ServeMux's Handle func, but it takes a flexwork
// Handler as its second parameter instead of a net/http Handler
func (m *MuxWrapper) Handle(pattern string, handler Handler) {
	m.m.Handle(pattern, Wrap(handler))
}

// HandleFunc is the equivalent to http.ServeMux's HandleFunc func, but it takes a flexwork
// Handler as its second parameter instead of a net/http Handler
func (m *MuxWrapper) HandleFunc(pattern string, handler HandlerFunc) {
	m.m.HandleFunc(pattern, WrapFunc(handler))
}

// ServeHTTP makes
func (m *MuxWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.m.ServeHTTP(w, r)
}
