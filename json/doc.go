// Package json is a convenience package for writing JSON APIs. Sample usage:
//
//  http.Handle(flexwork.WrapFunc(func(r *http.Request) Response {
//    // ...
//    json.Encode(http.StatusOK, someData)
//  }))
package json
