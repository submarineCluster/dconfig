// Package file is a file source. Expected format is json
package file

import (
	"io"
	"os"

	"git.code.oa.com/tencent_abtest/go-common/dconfig/source"
)

type file struct {
	path string
	opts source.Options
}

var (
	// DefaultPath TODO
	DefaultPath            = "config.json"
	envKeyServerConfigPath = "AB_SERVER_CONFIG_PATH"
)

// Read TODO
func (f *file) Read() (*source.ChangeSet, error) {
	fh, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	b, err := io.ReadAll(fh)
	if err != nil {
		return nil, err
	}
	info, err := fh.Stat()
	if err != nil {
		return nil, err
	}

	cs := &source.ChangeSet{
		Format:    format(f.path, f.opts.Encoder),
		Source:    f.String(),
		Timestamp: info.ModTime(),
		Data:      b,
	}
	cs.Checksum = cs.Sum()

	return cs, nil
}

// String TODO
func (f *file) String() string {
	return "file"
}

// Watch TODO
func (f *file) Watch() (source.Watcher, error) {
	if _, err := os.Stat(f.path); err != nil {
		return nil, err
	}
	return newWatcher(f)
}

// Write TODO
func (f *file) Write(cs *source.ChangeSet) error {
	return nil
}

// NewSource TODO
func NewSource(opts ...source.Option) source.Source {
	options := source.NewOptions(opts...)
	path := os.Getenv(envKeyServerConfigPath)
	if len(path) == 0 {
		path = DefaultPath
	}
	f, ok := options.Context.Value(filePathKey{}).(string)
	if ok {
		path = f
	}
	return &file{opts: options, path: path}
}
