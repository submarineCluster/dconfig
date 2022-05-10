package config

import (
	"time"

	"git.code.oa.com/tencent_abtest/go-common/dconfig/reader"
)

type value struct{}

func newValue() reader.Value {
	return new(value)
}

// Bool TODO
func (v *value) Bool(def bool) bool {
	return false
}

// Int TODO
func (v *value) Int(def int) int {
	return 0
}

// String TODO
func (v *value) String(def string) string {
	return ""
}

// Float64 TODO
func (v *value) Float64(def float64) float64 {
	return 0.0
}

// Duration TODO
func (v *value) Duration(def time.Duration) time.Duration {
	return time.Duration(0)
}

// StringSlice TODO
func (v *value) StringSlice(def []string) []string {
	return nil
}

// StringMap TODO
func (v *value) StringMap(def map[string]string) map[string]string {
	return map[string]string{}
}

// Scan TODO
func (v *value) Scan(val interface{}) error {
	return nil
}

// Bytes TODO
func (v *value) Bytes() []byte {
	return nil
}
