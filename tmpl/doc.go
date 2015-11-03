// Package tmpl contains a template rendering interface and some implementations
// and a convenience function to construct a flexwork.Response that renders a given
// template with some data. Sample usage:
//
//  http.HandleFunc(flexwork.WrapFunc(func(r *http.Request) Response {
//    // ...
//    tmpl.Render(http.StatusOK, renderer, "some_template", someData, "default_layout")
//  }))
package tmpl
