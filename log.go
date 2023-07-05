package rlog

import (
	"fmt"
	"strings"
)

//ModuleLogSwitchFunc 模块开关方法，如果返回true则打印日志否则不打印
type ModuleLogSwitchFunc func(moduleName string) bool

var fModuleSwitch ModuleLogSwitchFunc

type Log struct {
	w          []writer
	suffix     string // 日志文件名称前缀
	level      Level  // 日志级别
	conf       Config //日志配置
	mode       Mode
	moduleName string //模块名称
}

func NewRLog(conf Config) (*Log, error) {
	log := &Log{
		mode:  conf.Mode,
		level: conf.Level,
		conf:  conf,
		w:     make([]writer, 0),
	}
	if strings.Contains(string(conf.Mode), string(ModeFile)) {
		w, err := newZapWriter(conf.Path, conf.Suffix, conf.Suffix, conf.Skip)
		if err != nil {
			return nil, err
		}
		log.w = append(log.w, w)
	}
	if strings.Contains(string(conf.Mode), string(ModeConsole)) {
		w, err := newZapConsoleWriter(conf.Skip)
		if err != nil {
			return nil, err
		}
		log.w = append(log.w, w)
	}
	if len(log.w) == 0 {
		return nil, fmt.Errorf("log mode err must be one of [console,file,console|file]")
	}
	return log, nil
}

func (r *Log) Debug(msg string, m ...interface{}) {
	if !r.checkLogPrint(LevelDebug) {
		return
	}
	if r.moduleName != "" {
		m = append(m, "module", r.moduleName)
	}
	for _, w := range r.w {
		w.Debug(msg, m...)
	}
}

func (r *Log) Info(msg string, m ...interface{}) {
	if !r.checkLogPrint(LevelInfo) {
		return
	}
	if r.moduleName != "" {
		m = append(m, "module", r.moduleName)
	}
	for _, w := range r.w {
		w.Info(msg, m...)
	}
}

func (r *Log) Warn(msg string, m ...interface{}) {
	if !r.checkLogPrint(LevelWarn) {
		return
	}
	if r.moduleName != "" {
		m = append(m, "module", r.moduleName)
	}
	for _, w := range r.w {
		w.Warn(msg, m...)
	}
}

func (r *Log) Error(msg string, m ...interface{}) {
	if !r.checkLogPrint(LevelError) {
		return
	}
	if r.moduleName != "" {
		m = append(m, "module", r.moduleName)
	}
	for _, w := range r.w {
		w.Error(msg, m...)
	}
}

func (r *Log) addModuleNameToLogAny(m ...interface{}) {
	if r.moduleName == "" {
		return
	}
	m = append(m, "module", r.moduleName)
}

func (r *Log) SetLevel(level Level) {
	r.level = level
}

func (r *Log) GetLevel() Level {
	return r.level
}

func (r *Log) copy() *Log {
	return &Log{
		w:      r.w,
		suffix: r.suffix,
		level:  r.level,
		conf:   r.conf,
		mode:   r.mode,
	}
}

func (r *Log) CopyWithModuleName(moduleName string) *Log {
	cp := r.copy()
	cp.moduleName = moduleName
	return cp
}

func SetModuleSwitchFunc(switchFunc ModuleLogSwitchFunc) {
	fModuleSwitch = switchFunc
}

func (r *Log) checkModuleSwitch() bool {
	if r.moduleName == "" || fModuleSwitch == nil {
		return false
	}
	return fModuleSwitch(r.moduleName)
}

//checkLogPrint 检查日志是否应该打印
func (r *Log) checkLogPrint(level Level) bool {
	isOpen := r.checkModuleSwitch()
	if !levelCompare(r.level, level) && !isOpen {
		return false
	}
	return true
}
