package jsonx

import "strconv"

// Int forces the given value to be an int during the JSON unmarshalling,
// even if float was given.
type Int int

func (i *Int) UnmarshalJSON(b []byte) error {
	// Parse string as float.
	// This approach will cover both cases.
	v, err := strconv.ParseFloat(string(b), 64)
	// Set the value
	*i = Int(v)
	// Return error
	return err
}
