package tpl

import (
	"io"
)

// Executor is an interface for any template that can render itself with the
// given data (second param) and write the results to the given writer.
// Any instance of (*html/template).Template or (*text/template).Template
// implements this interface
type Executor interface {
	Execute(io.Writer, interface{}) error
}
