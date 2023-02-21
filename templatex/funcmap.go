//nolint:gochecknoglobals
package templatex

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/kyoto-framework/zen/v3/async"
	"github.com/kyoto-framework/zen/v3/b64"
	"github.com/kyoto-framework/zen/v3/conv"
	"github.com/kyoto-framework/zen/v3/errorsx"
	"github.com/kyoto-framework/zen/v3/fmtx"
	"github.com/kyoto-framework/zen/v3/jsonx"
	"github.com/kyoto-framework/zen/v3/mapx"
	"github.com/kyoto-framework/zen/v3/mathx"
	"github.com/kyoto-framework/zen/v3/regexpx"
	"github.com/kyoto-framework/zen/v3/slice"
)

/*
NestFuncMap allows to pack a given funcmap
into getter function.
In that way, you're able to use nested funcmap.

Usage:

	// Example of nesting
	example := template.FuncMap{ "jsonx": NestFuncMap(FuncMapJsonx) }
	// Example of usage
	"{{ jsonx.string .Value }}"
*/
func NestFuncMap(funcmap template.FuncMap) func() template.FuncMap {
	return func() template.FuncMap { return funcmap }
}

// FuncMap is a map of utilities functions to be used in templates.
var FuncMap = template.FuncMap{
	"mathx":   NestFuncMap(FuncMapmathx),
	"async":   NestFuncMap(FuncMapAsync),
	"b64":     NestFuncMap(FuncMapB64),
	"conv":    NestFuncMap(FuncMapConv),
	"errorsx": NestFuncMap(FuncMapErrorsx),
	"fmtx":    NestFuncMap(FuncMapFmtx),
	"jsonx":   NestFuncMap(FuncMapJsonx),
	"mapx":    NestFuncMap(FuncMapMapx),
	"regexpx": NestFuncMap(FuncMapRegexpx),
	"slice":   NestFuncMap(FuncMapSlice),

	"time":    NestFuncMap(FuncMapTime),
	"env":     NestFuncMap(FuncMapEnv),
	"strings": NestFuncMap(FuncMapStrings),
}

var FuncMapmathx = template.FuncMap{
	"div": mathx.DivRuntime,
	"mul": mathx.MulRuntime,
	"sub": mathx.SubRuntime,
	"sum": mathx.SumRuntime,
}

var FuncMapAsync = template.FuncMap{
	"await": async.AwaitRuntime,
}

var FuncMapB64 = template.FuncMap{
	"base64":      b64.Base64[string],
	"base64bytes": b64.Base64[[]byte],
	"string":      b64.String,
}

var FuncMapConv = template.FuncMap{
	"bool":    conv.Bool,
	"compose": conv.Compose,
	"float64": conv.Float64,
	"float":   conv.Float64,
	"int":     conv.Int,
	"ptr":     conv.PtrRuntime,
	"string":  conv.String,
}

var FuncMapErrorsx = template.FuncMap{
	"ignore": errorsx.IgnoreRuntime,
	"must":   errorsx.MustRuntime,
}

var FuncMapFmtx = template.FuncMap{
	"number":    fmtx.Number,
	"numberP0":  fmtx.NumberP0,
	"numberP1":  fmtx.NumberP1,
	"numeral":   fmtx.Numeral,
	"numeralP0": fmtx.NumeralP0,
	"numeralP1": fmtx.NumeralP1,
}

var FuncMapJsonx = template.FuncMap{
	"string": jsonx.String,
}

var FuncMapMapx = template.FuncMap{
	"compose": mapx.Compose,
}

var FuncMapRegexpx = template.FuncMap{
	"replace": regexpx.Replace,
}

var FuncMapSlice = template.FuncMap{
	"in":       slice.InRuntime,
	"newrange": slice.Range,
}

// Non-zen

var FuncMapTime = template.FuncMap{
	"now":    time.Now,
	"nowfmt": func(format string) string { return time.Now().Format(format) },
}

var FuncMapEnv = template.FuncMap{
	"env": os.Getenv,
}

var FuncMapStrings = template.FuncMap{
	"sprintf":  fmt.Sprintf,
	"contains": strings.Contains,
	"replace":  strings.ReplaceAll,
	"lower":    strings.ToLower,
	"upper":    strings.ToUpper,
	"title":    strings.Title, //nolint:staticcheck // Need to replace with something
	"trim":     strings.TrimSpace,
}
