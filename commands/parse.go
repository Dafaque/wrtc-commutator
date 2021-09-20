package commands

import (
	"bytes"
	"fmt"
)

const (
	SCOPE_BEGIN rune = 91
	SCOPE_END   rune = 93
)

func parseArg(kw string, src *[]byte) (value string, err error) {

	var token_found bool = false

	var scopeBeginCount int = 0
	var scopeEndCount int = 0

	var argBeginIdx int = 0
	var srcString = string(*src)
	for idx, r := range srcString {
		if r == SCOPE_BEGIN {
			scopeBeginCount++
			continue
		}
		if r == SCOPE_END {
			scopeEndCount++
			if scopeBeginCount == scopeEndCount {
				*src = append((*src)[:argBeginIdx], (*src)[idx+1:]...)
				*src = bytes.Trim(*src, " ")
				return
			}
			continue
		}

		value += string(r)

		if !token_found {
			if value == kw {
				value = ""
				argBeginIdx = idx + 1 - len(kw)
			}
			continue
		}
	}

	err = fmt.Errorf("syntax error")
	value = ""
	return
}

// for tests
var ForTestingOnly = parseArg
