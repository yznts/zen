package zen

import (
	"fmt"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatNumber(number float64, precision int, prefix, suffix string) string {
	p := strconv.Itoa(precision)
	v := message.NewPrinter(language.English).Sprintf("%."+p+"f", number)
	return prefix + v + suffix
}

func FormatNumberP0(number float64) string {
	return FormatNumber(number, 0, "", "")
}

func FormatNumberP1(number float64) string {
	return FormatNumber(number, 1, "", "")
}

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

func FormatNumberNumeralP0(number float64) string {
	return FormatNumberNumeral(number, 0)
}

func FormatNumberNumeralP1(number float64) string {
	return FormatNumberNumeral(number, 1)
}
