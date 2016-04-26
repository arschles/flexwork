package tpl

import (
	"html/template"
	"io"
	"text/template"
)

type Executor interface {
	Execute(io.Writer, interface{}) error
}

type Context struct {
	tplExecutor
}
