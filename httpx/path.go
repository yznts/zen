package httpx

import "strings"

/*
PathWrapper type is a wrapper around path string.
It provides a few useful extra methods to operate with path.
*/
type PathWrapper string

/*
Get returns path token at the given index.

Usage:

	path := httpx.Path("/foo/bar/baz")
	path.Get(1) // string{"bar"}
*/
func (p PathWrapper) Get(index int) string {
	tokens := p.Tokens()
	if len(tokens) < index {
		return ""
	}

	return tokens[index]
}

/*
GetBefore returns path token located before provided token.

Usage:

	path := httpx.Path("/foo/bar/baz")
	path.GetBefore("baz") // string{"bar"}
*/
func (p PathWrapper) GetBefore(token string) string {
	v, _ := p.GetBeforeWithIndex(token)

	return v
}

/*
GetBeforeWithIndex returns path token and it's index located before provided token.

Usage:

	path := httpx.Path("/foo/bar/baz")
	path.GetBeforeWithIndex("baz") // string{"bar"}, 1
*/
func (p PathWrapper) GetBeforeWithIndex(token string) (string, int) {
	tokens := p.Tokens()
	for i, ptoken := range tokens {
		if i == 0 {
			continue
		}

		if ptoken == token {
			return tokens[i-1], i - 1
		}
	}

	return "", -1
}

/*
GetAfter returns path token located after provided token.

Usage:

	path := httpx.Path("/foo/bar/baz")
	path.GetAfter("bar") // string{"baz"}
*/
func (p PathWrapper) GetAfter(token string) string {
	v, _ := p.GetAfterWithIndex(token)

	return v
}

/*
GetAfterWithIndex returns path token and it's index located after provided token.

Usage:

	path := httpx.Path("/foo/bar/baz")
	path.GetAfterWithIndex("bar") // string{"baz"}, 2
*/
func (p PathWrapper) GetAfterWithIndex(token string) (string, int) {
	tokens := p.Tokens()
	for i, ptoken := range tokens {
		if i+1 == len(tokens) {
			continue
		}

		if ptoken == token {
			return tokens[i+1], i + 1
		}
	}

	return "", -1
}

/*
Tokens returns path tokens.

Usage:

	path := httpx.Path("/foo/bar/baz")
	path.Tokens() // []string{"foo", "bar", "baz"}
*/
func (p PathWrapper) Tokens() []string {
	return strings.Split(string(p), "/")[1:]
}

/*
Path wraps a given path string into PathWrapper
to provide extra methods.

Usage:

	path := httpx.Path("/foo/bar/baz")
*/
func Path(path string) PathWrapper {
	return PathWrapper(path)
}

/*
Path wraps a given path tokens into PathWrapper
to provide extra methods.

Usage:

	path := httpx.PathFromTokens([]string{"foo", "bar", "baz"})
*/
func PathFromTokens(tokens []string) PathWrapper {
	return PathWrapper("/" + strings.Join(tokens, "/"))
}
