package tmpl

import (
	"net/http"
	"sync"
)

// Render is the result of a call to RenderRecorder's Render func
type Render struct {
	Code int
	Data interface{}
}

// RenderRecorder is an implementation of Renderer that records all calls to Render.
type RenderRecorder struct {
	// m synchronizes outputs
	m sync.RWMutex
	// Outputs stores the templates that have been rendered along with details on
	// each render
	renders map[string][]*Render
}

// NewRenderRecorder creates a new, empty RenderRecorder
func NewRenderRecorder() *RenderRecorder {
	return &RenderRecorder{m: sync.RWMutex{}, renders: make(map[string][]*Render)}
}

// Render is the interface implementation. Note that this func will call w.WriteHeader(code)
// as expected, but will write a string describing the template name that it would have rendered to w.
// call GetRenders(name) to get the actual rendered data.
func (t *RenderRecorder) Render(w http.ResponseWriter, code int, name string, data interface{}) {
	t.m.Lock()
	defer t.m.Unlock()
	rndrs, ok := t.renders[name]
	if !ok {
		rndrs = []*Render{}
	}
	rndrs = append(rndrs, &Render{Code: code, Data: data})
	t.renders[name] = rndrs
	w.WriteHeader(code)
	w.Write([]byte("would have rendered template" + name))
}

// GetRenders returns all of the renders that have been made for the given
// template name. returns an empty slice if none have ever been made
func (t *RenderRecorder) GetRenders(name string) []*Render {
	t.m.RLock()
	defer t.m.RUnlock()
	return t.renders[name]
}

// GetAllRenders returns all of the renders for each template name made
// to date.
func (t *RenderRecorder) GetAllRenders() map[string][]*Render {
	t.m.RLock()
	defer r.m.RUnlock()
	return t.renders
}
