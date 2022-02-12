//go:build !web

package model

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type SDP struct {
	From      []byte `msgpack:"from"`
	With      []byte `msgpack:"with"`
	As        byte   `msgpack:"as"`
	signature []byte
}

func (m *SDP) Verify(target []byte) bool {
	h := hmac.New(sha256.New, SECRET)
	h.Write(m.From)
	h.Write(m.With)
	h.Write(target)
	var sum []byte = h.Sum(nil)
	println(
		"signature:\ng:", hex.EncodeToString(m.signature),
		"\nc:", hex.EncodeToString(sum),
	)
	return bytes.EqualFold(m.signature, sum)
}

func NewSDP(from []byte, payload []byte, as byte, signature []byte) *SDP {
	return &SDP{
		signature: signature,
		From:      from,
		With:      payload,
		As:        as,
	}
}
