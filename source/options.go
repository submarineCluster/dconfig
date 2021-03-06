package source

import (
	"context"

	"git.code.oa.com/tencent_abtest/go-common/dconfig/encoder"
	"git.code.oa.com/tencent_abtest/go-common/dconfig/encoder/yaml"
	"go-micro.dev/v4/client"
)

// Options TODO
type Options struct {
	// Encoder
	Encoder encoder.Encoder

	// for alternative data
	Context context.Context

	// Client to use for RPC
	Client client.Client
}

// Option TODO
type Option func(o *Options)

// NewOptions TODO
func NewOptions(opts ...Option) Options {
	options := Options{
		Encoder: yaml.NewEncoder(),
		Context: context.Background(),
		Client:  client.DefaultClient,
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

// WithEncoder sets the source encoder
func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		o.Encoder = e
	}
}

// WithClient sets the source client
func WithClient(c client.Client) Option {
	return func(o *Options) {
		o.Client = c
	}
}
