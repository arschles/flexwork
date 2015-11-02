package tmpl

import (
	"net/http"

	"github.com/arschles/flexwork"
)

// Data is a convenience type for encodable template data
type Data interface{}

// Render is a convenience func that returns a flexwork.Response that will call renderer.Render
func Render(code int, renderer Renderer, name string, data Data, layout string) flexwork.Response {
	return ResponseFunc(func(w http.ResponseWriter) {
		renderer.Render(w, code, name, data, layout)
	})
}
