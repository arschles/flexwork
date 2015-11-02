package tmpl

import (
	"net/http"

	"github.com/unrolled/render"
)

// RenderRenderer is an implementation of Renderer that uses github.com/unrolled/render
// to do rendering
type RenderRenderer struct {
	rndr *render.Render
}

// NewRenderRenderer creates a new RenderRenderer
func NewRenderRenderer(r *render.Render) *RenderRenderer {
	return &RenderRenderer{rndr: r}
}

// Render is the interface implementation
func (r *RenderRenderer) Render(w http.ResponseWriter, code int, name string, data interface{}, layout string) {
	r.rndr.HTML(w, code, name, data, render.HTMLOptions{Layout: layout})
}
