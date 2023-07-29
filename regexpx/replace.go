package regexpx

import "regexp"

/*
Replace is a replace function similar to strings.ReplaceAll,
but with regex support.
*/
func ReplaceAll(s string, old string, new string) string { //nolint:predeclared
	r := regexp.MustCompile(old)

	return string(r.ReplaceAll([]byte(s), []byte(new)))
}
