/*
templatex - is a package that helps with templating.
It provides a way to use structs as funcmap extensions
and provides a set of extensions based on zen library.

You have at least 2 ways of using templatex:
using prepared funcmap, or using extensions separately.

To use prepared funcmap, just use predefined FuncMap:

	// Use templatex.FuncMap as the only source of functions
	funcmap := templatex.FuncMap
	// Combine with your own functions
	funcmap := mapx.Merge(
		templatex.FuncMap,
		ownfuncmap,
		...
	)

To use extensions separately, use import function and provide extension instance:

	funcmap := template.FuncMap{
		"jsonx": templatex.Import(&templatex.JsonExtension{}),
	}

In the template, you'll be able to use functions from the extension:

	{{ jsonx.String .Value }}

Each extension also provides own funcmap if you want to use methods as top-level,
instead of using dot notation:

	funcmap := mapx.Merge(
		templatex.Import(&ConvExtension{}).FuncMap(), // This will add "bool", "float", etc.
	)

In the template, you'll be able to use functions in this way:

	{{ bool .Value }}

Please, use index to explore available extensions.
*/
package templatex
