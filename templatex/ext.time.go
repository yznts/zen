package templatex

import (
	"html/template"
	"time"
)

/*
TimeExtension is a template compatibility wrapper
around existing time package.
*/
type TimeExtension struct{}

func (e *TimeExtension) Now() time.Time {
	return time.Now()
}

func (e *TimeExtension) Nowfmt(format string) string {
	return time.Now().Format(format)
}

func (e *TimeExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"now":    e.Now,
		"nowfmt": e.Nowfmt,
	}
}
