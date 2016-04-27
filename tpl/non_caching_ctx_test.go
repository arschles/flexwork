package tpl

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	numConcurrent = 10
)

func getBaseDir() string {
	gp := os.Getenv("GOPATH")
	return filepath.Join(gp, "src", "github.com", "arschles", "flexwork", "testdata")
}

func TestNonCachingCtxPrepare(t *testing.T) {
	ctx := NewNonCachingContext(getBaseDir())
	for i := 0; i < numConcurrent; i++ {
		go func(i int) {
			_, err := ctx.Prepare(NewFiles("tpl1.html", "tpl2.html"))
			if err != nil {
				t.Errorf("failed to prepare executor on run %d", i)
				return
			}
		}(i)
	}
}
