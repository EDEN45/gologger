package file

import (
	"fmt"
	"io"
)

const defaultErrorMessage = `{"message": "Ошибка записи лога", "err": "%s", "name": "%s"}`

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

func (s *Write) Write(str []byte) {
	if _, err := s.writer.Write(str); err != nil {
		fmt.Printf(defaultErrorMessage, err.Error(), s.writer.Name())
	}
}
