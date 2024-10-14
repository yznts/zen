package templatex

import (
	"html/template"

	"github.com/yznts/zen/v3/async"
)

/*
AsyncExtension is a template compatibility wrapper
around existing async zen package.
*/
type AsyncExtension struct{}

func (e *AsyncExtension) Await(f async.ImplementsAwaitRuntime) (any, error) {
	return async.AwaitRuntime(f)
}

func (e *AsyncExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"await": e.Await,
	}
}
