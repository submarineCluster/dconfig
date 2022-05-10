package config

import (
	"git.code.oa.com/tencent_abtest/go-common/dconfig/loader"
	"git.code.oa.com/tencent_abtest/go-common/dconfig/reader"
	"git.code.oa.com/tencent_abtest/go-common/dconfig/source"
)

// WithLoader sets the loader for manager config
func WithLoader(l loader.Loader) Option {
	return func(o *Options) {
		o.Loader = l
	}
}

// WithSource appends a source to list of sources
func WithSource(s source.Source) Option {
	return func(o *Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) Option {
	return func(o *Options) {
		o.Reader = r
	}
}

// WithWatcherDisabled TODO
func WithWatcherDisabled() Option {
	return func(o *Options) {
		o.WithWatcherDisabled = true
	}
}
