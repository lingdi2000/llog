package rlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewFileLog() (*Log, error) {
	return NewRLog(Config{
		Path:   "./test",
		Level:  LevelDebug,
		Mode:   ModeFile,
		Skip:   2,
		Suffix: "gamesrv",
	})
}

func NewConsoleLog() (*Log, error) {
	return NewRLog(Config{
		Level:  LevelDebug,
		Mode:   ModeConsole,
		Skip:   2,
		Suffix: "gamesrv",
	})
}

func TestLogLevel(t *testing.T) {
	fLog, err := NewFileLog()
	assert.Nil(t, err)
	fLog.Debug("this is a test", "data", "debug")
	fLog.Warn("this is a test", "data", "warn")
	fLog.Info("this is a test", "data", "info")
	fLog.Error("this is a test", "data", "error")

	cLog, err := NewRLog(Config{
		Level: LevelDebug,
		Mode:  ModeConsole,
		Skip:  2,
	})
	assert.Nil(t, err)
	cLog.Debug("this is a test", "data", "debug")
	cLog.Warn("this is a test", "data", "warn")
	cLog.Info("this is a test", "data", "info")
	cLog.Error("this is a test", "data", "error")

	cLog2, err := NewConsoleLog()
	assert.Nil(t, err)
	cLog2.Debug("this is a test", "data", "debug")
	cLog2.Warn("this is a test", "data", "warn")
	cLog2.Info("this is a test", "data", "info")
	cLog2.Error("this is a test", "data", "error")
}

func TestModuleName(t *testing.T) {
	cLog, err := NewConsoleLog()
	assert.Nil(t, err)
	moduleLog := cLog.CopyWithModuleName("payment")
	moduleLog.Info("pay info",
		"orderID", "F123D",
		"goods", "diamond",
	)
}

func TestOddItem(t *testing.T) {
	cLog, err := NewConsoleLog()
	assert.Nil(t, err)
	cLog.Info("odd example", "1", "2", 3)
}

func TestEvenItem(t *testing.T) {
	cLog, err := NewConsoleLog()
	assert.Nil(t, err)
	cLog.Info("even example", "1", "2",
		"3", "4",
		"5", "6",
	)
}

func TestModule(t *testing.T) {
	cLog, err := NewConsoleLog()
	cLog.level = LevelInfo
	assert.Nil(t, err)
	SetModuleSwitchFunc(func(moduleName string) bool {
		if moduleName == "login" {
			return true
		}
		return false
	})
	cLog.CopyWithModuleName("payment").Debug("pay info", "data", "orderInfo")
	cLog.CopyWithModuleName("login").Debug("login process", "data", "loginInfo")
}
