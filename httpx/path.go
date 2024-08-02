package httpx

import (
	"strings"
)

/*
Path type is a wrapper around path string.
It provides a few useful extra methods to operate with path.
*/
type Path string

/*
Get returns path token at the given index.

Usage:

	path := httpx.Path("/foo/bar/baz")
	path.Get(1) // string{"bar"}
*/
func (p Path) Get(index int) string {
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
func (p Path) GetBefore(token string) string {
	v, _ := p.GetBeforeWithIndex(token)

	return v
}

/*
GetBeforeWithIndex returns path token and it's index located before provided token.

Usage:

	path := httpx.Path("/foo/bar/baz")
	path.GetBeforeWithIndex("baz") // string{"bar"}, 1
*/
func (p Path) GetBeforeWithIndex(token string) (string, int) {
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
func (p Path) GetAfter(token string) string {
	v, _ := p.GetAfterWithIndex(token)

	return v
}

/*
GetAfterWithIndex returns path token and it's index located after provided token.

Usage:

	path := httpx.Path("/foo/bar/baz")
	path.GetAfterWithIndex("bar") // string{"baz"}, 2
*/
func (p Path) GetAfterWithIndex(token string) (string, int) {
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
func (p Path) Tokens() []string {
	return strings.Split(string(p), "/")[1:]
}

/*
PathFromTokens joins path tokens into a path string
and wraps it into Path type.

Usage:

	path := httpx.PathFromTokens([]string{"foo", "bar", "baz"})
*/
func PathFromTokens(tokens []string) Path {
	return Path("/" + strings.Join(tokens, "/"))
}
