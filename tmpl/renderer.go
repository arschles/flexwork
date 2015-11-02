package tmpl

import "net/http"

// Renderer is a generic template renderer.
type Renderer interface {
	// Render renders a template with the given data. This func should render the template identified
	// by layout if it's non-empty, and render the name template inside layout's {{yield}} function.
	// otherwise, this func should just render the name template.
	//
	// Before writing template results to w, this func should call w.WriteHeader(code)
	Render(w http.ResponseWriter, code int, name string, data interface{}, layout string)
}
