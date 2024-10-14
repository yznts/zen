package templatex

import (
	"html/template"

	"github.com/yznts/zen/v3/conv"
)

/*
ConvExtension is a template compatibility wrapper
around existing conv zen package.
*/
type ConvExtension struct{}

func (e *ConvExtension) Bool(val any) bool       { return conv.Bool(val) }
func (e *ConvExtension) Float(val any) float64   { return conv.Float64(val) }
func (e *ConvExtension) Float64(val any) float64 { return conv.Float64(val) }
func (e *ConvExtension) Int(val any) int         { return conv.Int(val) }
func (e *ConvExtension) Ptr(val any) any         { return conv.PtrRuntime(val) }
func (e *ConvExtension) String(val any) string   { return conv.String(val) }

func (e *ConvExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"bool":    e.Bool,
		"float":   e.Float,
		"float64": e.Float64,
		"int":     e.Int,
		"ptr":     e.Ptr,
		"string":  e.String,
	}
}
