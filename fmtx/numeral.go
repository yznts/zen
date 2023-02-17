package fmtx

import (
	"fmt"

	"github.com/kyoto-framework/zen/v3/conv"
)

/*
Numeral returns a shorten, string representation of the given number.

Usage:

	fmtx.Numeral(12345.456, 0) // "12k"
*/
func Numeral(number float64, precision int) string {
	p := conv.String(precision)
	switch {
	case number < 1000:
		return fmt.Sprintf("%.0f", number)
	case number >= 1000000000:
		return fmt.Sprintf("%."+p+"fB", number/1000000000)
	case number >= 1000000:
		return fmt.Sprintf("%."+p+"fM", number/1000000)
	case number >= 10000:
		return fmt.Sprintf("%.0fK", number/1000)
	case number >= 1000:
		return fmt.Sprintf("%."+p+"fK", number/1000)
	default:
		return fmt.Sprintf("%f", number)
	}
}

/*
NumeralP0 is a wrapper around FormatNumberNumeral with 0 precision.

Usage:

	fmtx.NumeralP0(12345.456) // "12k"
*/
func NumeralP0(number float64) string {
	return Numeral(number, 0)
}

/*
NumeralP1 is a wrapper around FormatNumberNumeral with 1 precision.

Usage:

	fmtx.NumeralP1(12345.456) // "12.3k"
*/
func NumeralP1(number float64) string {
	return Numeral(number, 1)
}
