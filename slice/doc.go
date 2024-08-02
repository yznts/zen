/*
slice - a package with common slice operations.
Allows filtering, selection, checking, modification, etc.

# Conditional

This section contains functions that check if a slice meets a certain condition.
Results are returned as a boolean value.

# Conditional - All

All checks if all elements in a slice meet a certain condition.

Usage:

	slice.All([]int{1, 2, 3}, func(v int) bool { return v > 0 }) // true
	slice.All([]int{1, 2, 3}, func(v int) bool { return v > 1 }) // false

# Conditional - Any

Any checks if any element in a slice meets a certain condition.

Usage:

	slice.Any([]int{1, 2, 3}, func(v int) bool { return v > 0 }) // true
	slice.Any([]int{1, 2, 3}, func(v int) bool { return v > 3 }) // false

# Conditional - Contains

Contains checks if a slice contains a certain value.

Usage:

	slice.Contains([]int{1, 2, 3}, 1) // true
	slice.Contains([]int{1, 2, 3}, 4) // false

# Transformation

This section contains functions that transform a slice into another slice.
Results are returned as a new slice.

# Transformation - Map

Map applies a function to each element in a slice and returns a new slice.

Usage:

	slice.Map([]int{1, 2, 3}, func(v int) int { return v * 2 }) // []int{2, 4, 6}

# Transformation - Filter

Filter applies a condition function to each element in a slice
and returns a new slice with only the elements that meet the condition.

Usage:

	slice.Filter([]int{1, 2, 3}, func(v int) bool { return v > 1 }) // []int{2, 3}

# Transformation - Chunks

Chunks splits a slice into smaller slices of a given size.

Usage:

	slice.Chunks([]int{1, 2, 3, 4, 5}, 2) // [[1 2] [3 4] [5]]

# Transformation - Cartesian

Cartesian returns a new slice with all possible combinations of elements from provided slices.

Usage:

	slice.Cartesian([]int{1, 2}, []int{3, 4}) // [[1 3] [1 4] [2 3] [2 4]]

# Transformation - Limit

Limit returns a new slice with the first n elements from the provided slice.

Usage:

	slice.Limit([]int{1, 2, 3}, 2) // []int{1, 2}

# Transformation - Unique

Unique returns a new slice with only unique elements from the provided slice.

Usage:

	slice.Unique([]int{1, 2, 2, 3}) // []int{1, 2, 3}

# Modification

This section contains functions to modify a slice.

# Modification - Insert

Insert injects a provided value into slice on the given index.

Usage:

	slice.Insert([]string{"b", "c"}, 0, "a") // []string{"a", "b", "c"}

# Selection

This section contains functions to select elements from a slice.

# Selection - First

First returns the first element in a slice.

Usage:

	slice.First([]int{1, 2, 3}) // 1

# Selection - Pop

Pop removes and returns the element from a slice.

Usage:

	slice, val := slice.Pop([]int{1, 2, 3}) // slice: []int{1, 2}, val: 3
	slice, val := slice.Pop([]int{1, 2, 3}, 0) // slice: []int{2, 3}, val: 1

# Other

This section contains functions that don't fit into any other category.

# Other - Index

Index returns the index of the first occurrence of a value in a slice.

Usage:

	slice.Index([]int{1, 2, 3}, 2) // 1

# Other - Count

Count returns the number of elements in a slice that meet a certain condition.

Usage:

	slice.Count([]int{1, 2, 3}, func(v int) bool { return v > 1 }) // 2

# Other - Range

Range returns a slice of integers from start to end (inclusive).

Usage:

	slice.Range(1, 5) // []int{1, 2, 3, 4, 5}
*/
package slice
