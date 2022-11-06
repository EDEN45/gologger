package gologger

import (
	"runtime"
	"strconv"
)

const skipTrace = 5

type traceLog struct {
	pc   uintptr
	file string
	line int
	ok   bool
}

func getTrace() *traceLog {
	v := &traceLog{}
	v.pc, v.file, v.line, v.ok = runtime.Caller(skipTrace)
	return v
}

func (t *traceLog) String() string {
	var fileName []rune

	rName := []rune(t.file)
	for i := len(rName) - 1; i >= 0; i-- {
		if rName[i] == '/' {
			break
		}
		fileName = append(fileName, rName[i])
	}

	var file = make([]rune, len(fileName))
	for i := len(fileName) - 1; i >= 0; i-- {
		file = append(file, fileName[i])
	}

	return string(file) + ":" + strconv.Itoa(t.line)
}
