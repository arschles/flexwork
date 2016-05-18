package tpl

import (
	"html/template"
)

// NewNonCachingContext creates and returns a new Context implementation
// that parses a new template on each call to prepare. This implementation
// is useful for development
func NewNonCachingContext(baseDir string) Context {
	return nonCachingContext{baseDir: baseDir}
}

type nonCachingContext struct {
	baseDir string
}

func (c nonCachingContext) Prepare(tplFiles Files) (Executor, error) {
	absPaths := tplFiles.absPaths(c.baseDir)
	ex, err := template.ParseFiles(absPaths...)
	if err != nil {
		return nil, err
	}
	return ex, nil
}
