package tpl

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/arschles/assert"
)

var (
	files = NewFiles("file1", "file2", "file3")
)

func TestFilesLen(t *testing.T) {
	assert.Equal(t, files.len(), len(files.list), "length")
}

func TestFilesMapKey(t *testing.T) {
	assert.Equal(t, files.mapKey(), strings.Join(files.list, ","), "map key")
}

func TestFilesAbsPaths(t *testing.T) {
	const absPath = "path1"
	absPaths := files.absPaths(absPath)
	assert.Equal(t, len(absPaths), files.len(), "slice length")
	for i, ap := range absPaths {
		assert.Equal(t, ap, filepath.Join(ap, files.list[i]), fmt.Sprintf("abs path %d", i))
	}
}
