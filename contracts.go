package gologger

type LevelLog int16

type Writer interface {
	Write(string2 string)
}
