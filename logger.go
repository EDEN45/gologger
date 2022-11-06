package gologger

import (
	"context"
	"time"
)

type Logger struct {
	in         chan []any
	done       chan bool
	appName    string
	timeFormat string
	writer     Writer
}

type Config struct {
	appName    string
	timeFormat string
	writer     Writer
}

func New(conf Config) *Logger {
	return &Logger{
		in:         make(chan []any, 10000),
		done:       make(chan bool, 1),
		appName:    conf.appName,
		timeFormat: conf.timeFormat,
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
				l.writer.Write(toJSON(v))
			}
		}
	}()
}

func (l *Logger) Wait() {
	<-l.done
}

func (l *Logger) Info(message string, vars []any) {
	l.log(LevelInfo, false, message, vars)
}

func (l *Logger) Debug(message string, vars []any) {
	l.log(LevelDebug, false, message, vars)
}

func (l *Logger) Warn(message string, vars []any) {
	l.log(LevelWarning, false, message, vars)
}

func (l *Logger) Error(message string, vars []any) {
	l.log(LevelError, false, message, vars)
}

func (l *Logger) Recover(message string, vars []any) {
	l.log(LevelRecover, true, message, vars)
}

func (l *Logger) log(lvl LevelLog, now bool, message string, vars []any) {
	values := l.combineAndMark(lvl, message, vars)
	if now {
		l.writer.Write(toJSON(values))
		return
	}

	l.in <- values
}

// combineAndMark объединяет все поля в слайс
func (l *Logger) combineAndMark(lvl LevelLog, message string, vars []any) []any {
	tr := getTrace()
	fields := []any{
		"time", time.Now().Format(l.timeFormat),
		"app", l.appName,
		"level", levelByCode(lvl),
		"message", message,
		"trace", tr.String(),
	}

	return append(fields, vars...)
}
