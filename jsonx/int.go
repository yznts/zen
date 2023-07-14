package jsonx

import "strconv"

// Int forces the given value to be an int during the JSON unmarshalling,
// even if float was given.
type Int int

func (i *Int) UnmarshalJSON(b []byte) error {
	// Handle null case.
	if string(b) == "null" {
		// Set default value
		*i = 0
		// Return nil
		return nil
	}
	// Parse string as float.
	// This approach will cover both int and float cases.
	v, err := strconv.ParseFloat(string(b), 64)
	// Set the value
	*i = Int(v)
	// Return error
	return err
}
