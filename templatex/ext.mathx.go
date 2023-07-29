package templatex

import (
	"html/template"

	"go.kyoto.codes/zen/v3/mathx"
)

/*
MathxExtension is a template compatibility wrapper
around existing mathx zen package.
*/
type MathxExtension struct{}

func (e *MathxExtension) DivRuntime(vals ...any) any { return mathx.DivRuntime(vals...) }
func (e *MathxExtension) MulRuntime(vals ...any) any { return mathx.MulRuntime(vals...) }
func (e *MathxExtension) SubRuntime(vals ...any) any { return mathx.SubRuntime(vals...) }
func (e *MathxExtension) SumRuntime(vals ...any) any { return mathx.SumRuntime(vals...) }

func (e *MathxExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"div": e.DivRuntime,
		"mul": e.MulRuntime,
		"sub": e.SubRuntime,
		"sum": e.SumRuntime,
	}
}
