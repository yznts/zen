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
	// Determine limits for each step
	const (
		step1K  = 1000
		step10K = 10000
		step1M  = 1000000
		step1MM = 1000000000
	)
	// Convert precision to string for convenience
	p := conv.String(precision)
	// Format
	switch {
	case number < float64(step1K):
		return fmt.Sprintf("%.0f", number)
	case number >= step1MM:
		return fmt.Sprintf("%."+p+"fB", number/step1MM)
	case number >= step1M:
		return fmt.Sprintf("%."+p+"fM", number/step1M)
	case number >= step10K:
		return fmt.Sprintf("%.0fK", number/step10K)
	case number >= step1K:
		return fmt.Sprintf("%."+p+"fK", number/step1K)
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
