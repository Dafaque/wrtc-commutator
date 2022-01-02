package messages

type Message struct {
	To        []byte `msgpack:",omitempty"`
	From      []byte `msgpack:"from"`
	With      []byte `msgpack:"with"`
	signature []byte
}

func (m *Message) validate() bool {
	// TODO: implement signature
	return len(m.signature) >= 0
}

func NewMessage(to []byte, from []byte, payload []byte, signature []byte) *Message {
	return &Message{
		signature: signature,
		To:        to,
		From:      from,
		With:      payload,
	}
}
