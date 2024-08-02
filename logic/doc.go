/*
logic - a package with common logic operations, like "or" or ternary.
Unfortunately, Go doesn't allow to do logic operations on values.

# Or / Tr

Or is a helper that returns the first non-empty value from the provided list.

Usage:

	// Acts like "||" in other languages
	val := logic.Or("", "default") // "default"
	val := logic.Or(0, 0, 1, 2) // 1

Tr acts like a ternary operator in other languages.
It returns one of two values based on the condition.

Usage:

	// Acts like "condition ? trueValue : falseValue" in other languages
	val := logic.Tr(false, "asd", "qwe") // "qwe"
	val := logic.Tr(true, "asd", "qwe") // "asd"

# Empty / NotEmpty

Empty is a helper that returns true if the provided value is the same as default for provided type.

Usage:

	// Acts like "bool(value)" in some languages
	val := logic.Empty("") // true
	val := logic.Empty(0) // true
	val := logic.Empty(false) // true
	val := logic.Empty(nil) // true
	val := logic.Empty(1) // false

NotEmpty is doing the opposite of Empty.

Usage:

	// Acts like "!bool(value)" in some languages
	val := logic.NotEmpty("") // false
	val := logic.NotEmpty(0) // false
	val := logic.NotEmpty(false) // false
	val := logic.NotEmpty(nil) // false
	val := logic.NotEmpty(1) // true
	val := logic.NotEmpty("asd") // true
	val := logic.NotEmpty(true) // true

	// This function can be combined with slice.Filter to remove empty values
	slice.Filter([]int{0, 0, 1, 4, 5}, logic.NotEmpty[int]) // []int{1, 4, 5}
*/
package logic
