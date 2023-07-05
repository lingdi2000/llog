package rlog

// Level defines the log level.
type Level int8

const (
	LevelDebug Level = iota
	LevelWarn
	LevelInfo
	LevelError
	LevelFatal
)

func levelCompare(src, dst Level) bool {
	return src <= dst
}

// Any 日志内容
type Any []interface{}

type Mode string

const (
	ModeConsole     Mode = "console"
	ModeFile        Mode = "file"
	ModeConsoleFile Mode = "console|file"
)
