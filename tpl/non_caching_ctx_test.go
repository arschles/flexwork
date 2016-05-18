package tpl

import (
	"testing"
)

func TestNonCachingCtxPrepare(t *testing.T) {
	errs := runCtxTest(NewNonCachingContext(getBaseDir()))
	if len(errs) > 0 {
		t.Fatalf("Errors: %s", errs)
	}
}
