//go:build web
package model

type SDP struct {
	From      string `json:"from"`
	With      string `json:"with"`
	As        string `json:"as"`
	signature string
}

func (m *SDP) validate() bool {
	return len(m.signature) >= 0
}

func NewSDP(from []byte, payload []byte, as byte, signature []byte) *SDP {
	return &SDP{
		signature: string(signature),
		From:      string(from),
		With:      string(payload),
		As:        string(as),
	}
}
