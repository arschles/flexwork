package tpl

import (
	"html/template"
	"sync"
)

// NewCachingContext creates and returns a new Context implementation
// that lazily parses a new template on the first call to Prepare, and
// caches that template on subsequent calls
func NewCachingContext(baseDir string) Context {
	return &cachingContext{
		baseDir: baseDir,
		mut:     new(sync.Mutex),
		cache:   make(map[filesMapKey]*template.Template),
	}
}

type cachingContext struct {
	baseDir string
	// TODO: speed cache filling up. use a map of cond vars and a background worker
	mut   *sync.Mutex
	cache map[filesMapKey]*template.Template
}

func (c *cachingContext) Prepare(tplFiles Files) (*template.Template, error) {
	c.mut.Lock()
	defer c.mut.Unlock()
	mk := tplFiles.mapKey()
	tpl, ok := c.cache[mk]
	if !ok {
		absPaths := tplFiles.absPaths(c.baseDir)
		t, err := template.ParseFiles(absPaths...)
		if err != nil {
			return nil, err
		}
		tpl = t
	}
	return tpl, nil
}
