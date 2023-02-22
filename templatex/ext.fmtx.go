package templatex

import (
	"html/template"

	"github.com/kyoto-framework/zen/v3/fmtx"
)

/*
AsyncExtension is a template compatibility wrapper
around existing async zen package.
*/
type FmtxExtension struct{}

func (e *FmtxExtension) Number(number float64, precision int, prefix, suffix string) string {
	return fmtx.Number(number, precision, prefix, suffix)
}

func (e *FmtxExtension) NumberP0(number float64) string {
	return fmtx.NumberP0(number)
}

func (e *FmtxExtension) NumberP1(number float64) string {
	return fmtx.NumberP1(number)
}

func (e *FmtxExtension) Numeral(number float64, precision int) string {
	return fmtx.Numeral(number, precision)
}

func (e *FmtxExtension) NumeralP0(number float64) string {
	return fmtx.NumeralP0(number)
}

func (e *FmtxExtension) NumeralP1(number float64) string {
	return fmtx.NumeralP1(number)
}

func (e *FmtxExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"number":    e.Number,
		"numberP0":  e.NumberP0,
		"numberP1":  e.NumberP1,
		"numeral":   e.Numeral,
		"numeralP0": e.NumeralP0,
		"numeralP1": e.NumeralP1,
	}
}
