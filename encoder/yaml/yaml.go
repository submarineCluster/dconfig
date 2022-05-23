// Package yaml TODO
package yaml

import (
	"git.code.oa.com/tencent_abtest/go-common/dconfig/encoder"
	"gopkg.in/yaml.v3"
)

type yamlEncoder struct{}

// Encode TODO
func (y yamlEncoder) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// Decode TODO
func (y yamlEncoder) Decode(d []byte, v interface{}) error {
	return yaml.Unmarshal(d, v)
}

// String TODO
func (y yamlEncoder) String() string {
	return "yaml"
}

// NewEncoder TODO
func NewEncoder() encoder.Encoder {
	return yamlEncoder{}
}
