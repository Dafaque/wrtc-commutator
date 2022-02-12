package commands

import (
	"bytes"
	"commutator/errcodes"
)

const (
	SCOPE_BEGIN byte = 91
	SCOPE_END   byte = 93
)

func parseArg(kw []byte, src *[]byte) (value []byte, err errcodes.ErrorCode) {

	var token_found bool = false

	var scopeBeginCount int = 0
	var scopeEndCount int = 0

	var argBeginIdx int = 0
	for idx, r := range *src {
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
		if value == nil {
			value = []byte{}
		}
		value = append(value, r)

		if !token_found {
			if bytes.EqualFold(value, kw) {
				value = nil
				argBeginIdx = idx + 1 - len(kw)
			}
			continue
		}
	}

	err = errcodes.ERROR_CODE_INVALID_SYNTAX
	value = nil
	return
}
