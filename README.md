
# Zen

Zen is an utilities package that provides a set of commonly used functions, but missing from the standard library (for now).
Most of the functions are adapted to be used with `html/template`.

## Motivation

Too much helper files that migrating from project to project with copy-paste approach.
Would be nice to include this functions as a part of kyoto project.

## Features

- Aggregative (Min, Max, Avg)
- Arithmetic (Sum, Sub, Mul, Div)
- Transform (Ptr, Bool, Int, Float64, String, JSON, B64Enc, B64Dec, Compose)
- Slice casting (CastSlice)
- Number formatting (FormatNumber, FormatNumberNumeral)
- Range functions (Filter, Map, Range, In)
- Templates specific functions (type casting, composition, string operations, etc)

## Documentation

Code is well-documented, so it's pretty comfortable to use documentation provided by [pkg.go.dev](https://pkg.go.dev/github.com/kyoto-framework/zen)
