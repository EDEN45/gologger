package file

import (
	"fmt"
	"io"
)

type NameWriter interface {
	io.Writer
	Name() string
}

type Write struct {
	writer NameWriter
}

func New(file NameWriter) *Write {
	return &Write{
		writer: file,
	}
}

func (s *Write) Write(str string) {
	if _, err := s.writer.Write([]byte(str)); err != nil {
		fmt.Printf(`{"message": "Ошибка записи лога", "err": "%s", "name": "%s"}`, err.Error(), s.writer.Name())
	}
}
