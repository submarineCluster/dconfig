package memory

import (
	"git.code.oa.com/tencent_abtest/go-common/dconfig/loader"
	"git.code.oa.com/tencent_abtest/go-common/dconfig/reader"
	"git.code.oa.com/tencent_abtest/go-common/dconfig/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}

// WithWatcherDisabled TODO
func WithWatcherDisabled() loader.Option {
	return func(o *loader.Options) {
		o.WithWatcherDisabled = true
	}
}
