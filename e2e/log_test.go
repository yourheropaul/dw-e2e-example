package e2e

import (
	"strings"
	"testing"
)

type writeCloser struct {
	t *testing.T
}

func (w *writeCloser) Write(p []byte) (n int, err error) {
	w.t.Log(strings.TrimRight(string(p), "\n"))
	return len(p), nil
}

func (w *writeCloser) Close() (err error) {
	return nil
}

func mockServiceLogger(t *testing.T, name string) *ServiceLogger {
	return NewLog().NewServiceLogger(name, &writeCloser{t: t})
}
