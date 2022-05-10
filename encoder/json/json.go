// Package json TODO
package json

import (
	"encoding/json"

	"git.code.oa.com/tencent_abtest/go-common/dconfig/encoder"
)

type jsonEncoder struct{}

// Encode TODO
func (j jsonEncoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Decode TODO
func (j jsonEncoder) Decode(d []byte, v interface{}) error {
	return json.Unmarshal(d, v)
}

// String TODO
func (j jsonEncoder) String() string {
	return "json"
}

// NewEncoder TODO
func NewEncoder() encoder.Encoder {
	return jsonEncoder{}
}
