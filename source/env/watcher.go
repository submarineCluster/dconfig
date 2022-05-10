package env

import (
	"git.code.oa.com/tencent_abtest/go-common/dconfig/source"
)

type watcher struct {
	exit chan struct{}
}

// Next TODO
func (w *watcher) Next() (*source.ChangeSet, error) {
	<-w.exit

	return nil, source.ErrWatcherStopped
}

// Stop TODO
func (w *watcher) Stop() error {
	close(w.exit)
	return nil
}

func newWatcher() (source.Watcher, error) {
	return &watcher{exit: make(chan struct{})}, nil
}
