package templatex

/*
Import allows to pack a given struct
into getter function.
In that way, you're able to use struct and struct methods as funcmap extensions.

Usage:

	// Example of import
	example := template.FuncMap{ "jsonx": Import(&templatex.JsonxExtension{}) }
	// Example of usage
	"{{ jsonx.String .Value }}"
*/
func Import[T any](extension T) func() T {
	return func() T { return extension }
}
