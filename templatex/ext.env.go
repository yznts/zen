package templatex

import (
	"html/template"
	"os"
)

/*
EnvExtension is a template compatibility wrapper
around existing environment-related package.
*/
type EnvExtension struct{}

func (e *EnvExtension) Get(key string) string {
	return os.Getenv(key)
}

func (e *EnvExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"env": e.Get,
	}
}
