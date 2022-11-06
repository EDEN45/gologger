package gologger

// Const of type for type notice
const (
	_ LevelLog = iota - 1
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelReport
	LevelRecover
)

var levelCodes = []string{
	LevelDebug:   "DEBUG",
	LevelInfo:    "INFO",
	LevelWarning: "WARN",
	LevelError:   "ERROR",
	LevelReport:  "REPORT",
	LevelRecover: "RECOVER",
}

func levelByCode(code LevelLog) string {
	return levelCodes[code]
}
