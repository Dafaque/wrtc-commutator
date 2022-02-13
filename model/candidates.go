//go:build !web

package model

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// TODO: make single message type
type Candidates struct {
	From      []byte `msgpack:"from"`
	With      []byte `msgpack:"with"`
	signature []byte
}

func (c *Candidates) Verify(target []byte) bool {
	h := hmac.New(sha256.New, SECRET)
	h.Write(c.From)
	h.Write(c.With)
	h.Write(target)
	var sum []byte = h.Sum(nil)
	var signature []byte = make([]byte, hex.DecodedLen(len(c.signature)))
	hex.Decode(signature, c.signature)
	return bytes.EqualFold(signature, sum)
}

func NewCandidates(from []byte, payload []byte, signature []byte) *Candidates {
	return &Candidates{
		signature: signature,
		From:      from,
		With:      payload,
	}
}
