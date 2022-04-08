package zen

import (
	"html/template"
	"os"
	"strings"
	"time"
)

// FuncMap is a map of functions to be used in templates.
func FuncMap() template.FuncMap {
	return template.FuncMap{
		// Cast
		"html":     func(val string) template.HTML { return template.HTML(val) },
		"htmlattr": func(val string) template.HTMLAttr { return template.HTMLAttr(val) },
		"url":      func(val string) template.URL { return template.URL(val) },
		"css":      func(val string) template.CSS { return template.CSS(val) },
		"js":       func(val string) template.JS { return template.JS(val) },
		// Transform
		"compose": Compose,
		"ptr":     func(val any) any { return &val },
		"bool":    Bool,
		"int":     Int,
		"float":   Float64,
		"string":  String,
		"json":    JSON,
		"b64enc":  B64Dec,
		"b64dec":  B64Enc,
		// Arithmetic
		"sum": SumRuntime,
		"sub": SubRuntime,
		"mul": MulRuntime,
		"div": DivRuntime,
		// Strings functions
		"contains": strings.Contains,
		"replace":  strings.ReplaceAll,
		"lower":    strings.ToLower,
		"upper":    strings.ToUpper,
		"title":    strings.Title,
		"trim":     strings.TrimSpace,
		// Numbers formatting
		"number":    FormatNumber,
		"numberP0":  FormatNumberP0,
		"numberP1":  FormatNumberP1,
		"numeral":   FormatNumberNumeral,
		"numeralP0": FormatNumberNumeralP0,
		"numeralP1": FormatNumberNumeralP1,
		// Time
		"now":  func() time.Time { return time.Now() },
		"date": func(format string) string { return time.Now().Format(format) },
		// Range
		"rng": Range,
		"in":  In,
		// Environment
		"env": os.Getenv,
	}
}
