//nolint:gochecknoglobals
package templatex

import (
	"html/template"
)

// FuncMap is a map of utilities functions to be used in templates.
var FuncMap = template.FuncMap{
	"mathx":   Import(&MathxExtension{}),
	"async":   Import(&AsyncExtension{}),
	"b64":     Import(&B64Extension{}),
	"conv":    Import(&ConvExtension{}),
	"errorsx": Import(&ErrorsxExtension{}),
	"fmtx":    Import(&FmtxExtension{}),
	"jsonx":   Import(&JsonxExtension{}),
	"mapx":    Import(&MapxExtension{}),
	"regexpx": Import(&RegexpxExtension{}),
	"slice":   Import(&SliceExtension{}),

	"time":    Import(&TimeExtension{}),
	"env":     Import(&EnvExtension{}),
	"strings": Import(&StringsExtension{}),
}
