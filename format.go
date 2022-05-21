/*
	-

	Format

	Zen provides some useful formatting functions.

	Examples:

		func main() {
			zen.FormatNumber(12345.456, 0, "$", "") // "$12,345"
			zen.FormatNumberP0(12345.456) // "12,345"
			zen.FormatNumberP1(12345.456) // "12,345.4"
			zen.FormatNumberNumeral(12345.456, 0) // "12k"
			zen.FormatNumberNumeralP0(12345.456) // "12k"
			zen.FormatNumberNumeralP1(12345.456) // "12.3k"
		}
*/
package zen

import (
	"fmt"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// FormatNumber returns a string representation of the given number in specified format.
//
// Usage:
//  zen.FormatNumber(12345.456, 0, "$", "") // "$12,345"
func FormatNumber(number float64, precision int, prefix, suffix string) string {
	p := strconv.Itoa(precision)
	v := message.NewPrinter(language.English).Sprintf("%."+p+"f", number)
	return prefix + v + suffix
}

// FormatNumberP0 is a wrapper around FormatNumber with 0 precision and no prefix or suffix.
//
// Usage:
//  zen.FormatNumberP0(12345.456) // "12,345"
func FormatNumberP0(number float64) string {
	return FormatNumber(number, 0, "", "")
}

// FormatNumberP1 is a wrapper around FormatNumber with 1 precision and no prefix or suffix.
//
// Usage:
//  zen.FormatNumberP1(12345.456) // "12,345.4"
func FormatNumberP1(number float64) string {
	return FormatNumber(number, 1, "", "")
}

// FormatNumberNumeral returns a shorten, string representation of the given number.
//
// Usage:
//  zen.FormatNumberNumeral(12345.456, 0) // "12k"
func FormatNumberNumeral(number float64, precision int) string {
	p := strconv.Itoa(precision)
	var s string
	if number < 1000 {
		s += fmt.Sprintf("%.0f", number)
	} else if number >= 1000000000 {
		s += fmt.Sprintf("%."+p+"fB", number/1000000000)
	} else if number >= 1000000 {
		s += fmt.Sprintf("%."+p+"fM", number/1000000)
	} else if number >= 10000 {
		s += fmt.Sprintf("%.0fK", number/1000)
	} else if number >= 1000 {
		s += fmt.Sprintf("%."+p+"fK", number/1000)
	}
	return s
}

// FormatNumberNumeralP0 is a wrapper around FormatNumberNumeral with 0 precision.
//
// Usage:
//  zen.FormatNumberNumeralP0(12345.456) // "12k"
func FormatNumberNumeralP0(number float64) string {
	return FormatNumberNumeral(number, 0)
}

// FormatNumberNumeralP1 is a wrapper around FormatNumberNumeral with 1 precision.
//
// Usage:
//  zen.FormatNumberNumeralP1(12345.456) // "12.3k"
func FormatNumberNumeralP1(number float64) string {
	return FormatNumberNumeral(number, 1)
}
