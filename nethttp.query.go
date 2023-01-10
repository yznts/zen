package zen

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// QueryWrapper type is a wrapper for url.Values.
// It provides a few useful extra methods.
type QueryWrapper struct {
	url.Values
}

// Unmarshal helps to parse url.Values into a struct.
// Slightly modified version of github.com/knadh/querytostruct
//
// Example:
//
//	var target struct {
//		Foo string `query:"foo"`
//		Bar int `query:"bar"`
//	}
//
//	q, _ := url.ParseQuery("foo=asdqwe&bar=123")
//	kyoto.Query(q).Unmarshal(&target)
func (q *QueryWrapper) Unmarshal(target any) error {
	// Get target reflection value
	ob := reflect.ValueOf(target)
	if ob.Kind() == reflect.Ptr {
		ob = ob.Elem()
	}

	// Validate value is a struct
	if ob.Kind() != reflect.Struct {
		return fmt.Errorf("failed to encode form values to struct, non struct type: %T", ob)
	}

	// Go through every field in the struct and look for it in the query map
	for i := 0; i < ob.NumField(); i++ {
		f := ob.Field(i)
		if f.IsValid() && f.CanSet() {
			tag := ob.Type().Field(i).Tag.Get("query")
			if tag == "" || tag == "-" {
				continue
			}

			// Got a struct field with a tag.
			// If that field exists in the arg and convert its type.
			// Tags are of the type `tagname,attribute`
			tag = strings.Split(tag, ",")[0]
			if _, ok := q.Values[tag]; !ok {
				continue
			}

			// The struct field is a slice type.
			if f.Kind() == reflect.Slice {
				var (
					vals    = q.Values[tag]
					numVals = len(vals)
				)

				// Make a slice.
				sl := reflect.MakeSlice(f.Type(), numVals, numVals)

				// If it's a []byte slice (=[]uint8), assign here.
				if f.Type().Elem().Kind() == reflect.Uint8 {

					br := q.Get(tag)
					b := make([]byte, len(br))
					copy(b, br)
					f.SetBytes(b)
					continue
				}

				// Iterate through args and assign values
				// to each item in the slice.
				for i, v := range vals {
					querySetVal(sl.Index(i), string(v))
				}
				f.Set(sl)
			} else {
				querySetVal(f, string(q.Get(tag)))
			}
		}
	}
	return nil
}

func querySetVal(f reflect.Value, val string) bool {
	switch f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v, err := strconv.ParseInt(val, 10, 0); err == nil {
			f.SetInt(v)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v, err := strconv.ParseUint(val, 10, 0); err == nil {
			f.SetUint(v)
		}
	case reflect.Float32:
		if v, err := strconv.ParseFloat(val, 32); err == nil {
			f.SetFloat(v)
		}
	case reflect.Float64:
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			f.SetFloat(v)
		}
	case reflect.String:
		f.SetString(val)
	case reflect.Bool:
		b, _ := strconv.ParseBool(val)
		f.SetBool(b)
	default:
		return false
	}
	return true
}

func Query(q url.Values) *QueryWrapper {
	return &QueryWrapper{q}
}
