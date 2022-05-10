package reader

import (
	"git.code.oa.com/tencent_abtest/go-common/dconfig/encoder"
	"git.code.oa.com/tencent_abtest/go-common/dconfig/encoder/json"
)

// Options TODO
type Options struct {
	Encoding map[string]encoder.Encoder
}

// Option TODO
type Option func(o *Options)

// NewOptions TODO
func NewOptions(opts ...Option) Options {
	options := Options{
		Encoding: map[string]encoder.Encoder{
			"json": json.NewEncoder(),
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

// WithEncoder TODO
func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		if o.Encoding == nil {
			o.Encoding = make(map[string]encoder.Encoder)
		}
		o.Encoding[e.String()] = e
	}
}
