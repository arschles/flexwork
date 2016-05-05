package tpl

import (
	"os"
	"path/filepath"
	"sync"
)

const (
	numConcurrent = 10
)

func getBaseDir() string {
	gp := os.Getenv("GOPATH")
	return filepath.Join(gp, "src", "github.com", "arschles", "flexwork", "testdata")
}

func runCtxTest(ctx Context) []error {
	var wg sync.WaitGroup
	errCh := make(chan error)
	for i := 0; i < numConcurrent; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, err := ctx.Prepare(NewFiles("tpl1.html", "tpl2.html"))
			if err != nil {
				errCh <- err
				return
			}
		}(i)
	}
	go func() {
		wg.Wait()
		close(errCh)
	}()
	errs := []error{}
	for {
		err, ok := <-errCh
		if !ok {
			break
		}
		errs = append(errs, err)
	}
	return errs
}
