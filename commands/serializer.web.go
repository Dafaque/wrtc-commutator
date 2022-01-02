//go:build web
package commands

import (
	"encoding/json"
)

var Serialize Serializer = func(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}
