/*
mapx - a package with common map functions,
like keys, values, merge, etc.

# Construction

This section contains functions that help with map construction,
like creating a map from a list of values, or merging maps.

# Construction - Compose

Compose creates a map from a list,
getting key/value one by one.

Usage:

	mapx.Compose("a", 1, "b", 2) // map[any]any{"a": 1, "b": 2}

# Construction - Merge

Merge merges multiple maps into one.

Usage:

	mapx.Merge(  // map[string]int{"a": 1, "b": 3, "c": 4}
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 3, "c": 4},
	)

# Transformation

This section contains functions that transform a map into another object.

# Transformation - Keys

Keys returns a list of keys from a map.

Usage:

	mapx.Keys(map[string]int{"a": 1, "b": 2}) // []string{"a", "b"}

# Transformation - Values

Values returns a list of values from a map.

Usage:

	mapx.Values(map[string]int{"a": 1, "b": 2}) // []int{1, 2}

# Modification

This section contains functions that modify an existing map.

# Modification - Keep

Keep keeps only keys that are in the list.

Usage:

	mapx.Keep(map[string]int{"a": 1, "b": 2}, "a") // map[string]int{"a": 1}

# Modification - Delete

Delete removes keys from a map.

Usage:

	mapx.Delete(map[string]int{"a": 1, "b": 2}, "a") // map[string]int{"b": 2}
*/
package mapx
