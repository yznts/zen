package templatex

import (
	"fmt"
	"html/template"
	"strings"
)

/*
StringsExtension is a template compatibility wrapper
around existing string-related packages.
*/
type StringsExtension struct{}

func (e *StringsExtension) Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func (e *StringsExtension) Contains(str string, substr string) bool {
	return strings.Contains(str, substr)
}

func (e *StringsExtension) Replace(str, old, new string) string {
	return strings.ReplaceAll(str, old, new)
}

func (e *StringsExtension) Lower(str string) string {
	return strings.ToLower(str)
}

func (e *StringsExtension) Upper(str string) string {
	return strings.ToUpper(str)
}

func (e *StringsExtension) Title(str string) string {
	return strings.Title(str)
}

func (e *StringsExtension) Trim(str string) string {
	return strings.TrimSpace(str)
}

func (e *StringsExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"sprintf":  e.Sprintf,
		"contains": e.Contains,
		"replace":  e.Replace,
		"lower":    e.Lower,
		"upper":    e.Upper,
		"title":    e.Title,
		"trim":     e.Trim,
	}
}
