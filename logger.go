package gologger

import (
	"context"
	"fmt"
)

type Writer interface {
	Write([]byte)
}

type Logger struct {
	in     chan []any
	done   chan bool
	writer Writer
}

type Config struct {
}

func New(conf Config) *Logger {
	return &Logger{
		in:   make(chan []any, 10000),
		done: make(chan bool, 1),
	}
}

func (l *Logger) Run(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				l.done <- true
				return
			case v := <-l.in:
				fmt.Println(v)
				l.writer.Write([]byte("example"))
			}
		}
	}()
}

func (l *Logger) Wait() {
	<-l.done
}

func (l *Logger) Log(vars []any) {
	l.in <- vars
}
