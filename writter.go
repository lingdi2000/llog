package rlog

//writer 定义一个日志writer方便替换底层日志库
type writer interface {
	Debug(msg string, m ...interface{})
	Info(msg string, m ...interface{})
	Warn(msg string, m ...interface{})
	Error(msg string, m ...interface{})
}
