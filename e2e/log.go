package e2e

import (
	"fmt"
	"io"
	"os"
)

type Log struct {
	combined []io.WriteCloser
	services map[string]*ServiceLogger
}

func NewLog(writers ...io.WriteCloser) *Log {
	return &Log{
		combined: writers,
		services: map[string]*ServiceLogger{},
	}
}

func (l *Log) NewServiceLogger(name string, writeCloser io.WriteCloser) *ServiceLogger {
	sl := &ServiceLogger{
		parent:      l,
		name:        name,
		writeCloser: writeCloser,
	}

	l.services[name] = sl

	return sl
}

func (l *Log) Write(p []byte) (n int, err error) {
	for _, service := range l.combined {
		if n, err := service.Write(p); err != nil {
			return n, err
		}
	}
	return len(p), nil
}

func (l *Log) Close() error {
	for _, service := range l.combined {
		if service == os.Stdout {
			continue
		}

		if err := service.Close(); err != nil {
			return err
		}
	}

	for _, service := range l.services {
		if err := service.Close(); err != nil {
			return err
		}
	}

	return nil
}

type ServiceLogger struct {
	parent      *Log
	name        string
	writeCloser io.WriteCloser
}

func (l *ServiceLogger) Write(p []byte) (n int, err error) {
	if n, err := l.writeCloser.Write(p); err != nil {
		return n, err
	}

	prefixed := fmt.Sprintf("[%s] %s", l.name, p)

	if n, err := l.parent.Write([]byte(prefixed)); err != nil {
		return n, err
	}
	return len(p), nil
}

func (l *ServiceLogger) Close() error {
	return l.writeCloser.Close()
}

func (l *ServiceLogger) Println(args ...interface{}) {
	l.Write([]byte(fmt.Sprint(args...) + "\n"))
}

func (l *ServiceLogger) Printf(format string, args ...interface{}) {
	l.Write([]byte(fmt.Sprintf(format, args...) + "\n"))
}
