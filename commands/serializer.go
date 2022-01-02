//go:build !web
package commands

import "github.com/vmihailenco/msgpack/v5"

var Serialize Serializer = func(i interface{}) ([]byte, error) {
	return msgpack.Marshal(i)
}
