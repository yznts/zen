package fmtx

import (
	"go.kyoto.codes/zen/v3/conv"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

/*
Number returns a string representation of the given number in specified format.

Usage:

	fmtx.Number(12345.456, 0, "$", "") // "$12,345"
*/
func Number(number float64, precision int, prefix, suffix string) string {
	var (
		precisionstr = conv.String(precision)
		numstr       = message.NewPrinter(language.English).Sprintf("%."+precisionstr+"f", number)
	)

	return prefix + numstr + suffix
}

/*
NumberP0 is a wrapper around FormatNumber with 0 precision and no prefix or suffix.

Usage:

	fmtx.NumberP0(12345.456) // "12,345"
*/
func NumberP0(number float64) string {
	return Number(number, 0, "", "")
}

/*
NumberP1 is a wrapper around FormatNumber with 1 precision and no prefix or suffix.

Usage:

	fmtx.NumberP1(12345.456) // "12,345.4"
*/
func NumberP1(number float64) string {
	return Number(number, 1, "", "")
}
