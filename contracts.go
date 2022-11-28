package gologger

type LevelLog int16

type Writer interface {
	Write(str []byte)
}
