package tpl

import (
	"path/filepath"
	"strings"
)

// Files is an ordered list of filenames that comprise a template.
// The first filename in the list should be the primary template that will
// be rendered, and the remaining ones should be supporting templates like
// block definitions
type Files struct {
	list      []string
	mapKeyStr string
}

type filesMapKey string

// NewFiles creates a new Files struct from the given filenames
func NewFiles(files ...string) *Files {
	return &Files{
		list:      files,
		mapKeyStr: strings.Join(files, ","),
	}
}

func (f Files) len() int {
	return len(f.list)
}

func (f *Files) absPaths(absPath string) []string {
	ret := make([]string, f.len())
	for i, fileName := range f.list {
		ret[i] = filepath.Join(absPath, fileName)
	}
	return ret
}

func (f Files) mapKey() filesMapKey {
	return filesMapKey(f.mapKeyStr)
}
