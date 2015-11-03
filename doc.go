// Package flexwork is a new take on Go web frameworks. It focuses on each individual handler, leaving the choice of
// mux, database, template renderer, etc... up to you. flexport handlers are different from net/http ones because they
// return a response like so:
//
//  type Handler interface {
//    ServeHTTP(r *http.Request) Response
//  }
//
// This change is significant because the framework (or you) can use each returned Response as a signal
// that the request is finished (or at least part of it has finished). More information to come here.
package flexwork
