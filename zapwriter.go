package rlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const callerSkipOffset = 3

type zapWriter struct {
	log *zap.Logger
}

func newZapConsoleWriter(skip int) (writer *zapWriter, err error) {
	if skip == 0 {
		skip = callerSkipOffset
	}
	zapEnc := getZapEncode()
	core := zapcore.NewCore(zapcore.NewJSONEncoder(zapEnc), zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(skip))
	writer = &zapWriter{
		log: logger,
	}
	return
}

func getZapEncode() zapcore.EncoderConfig {
	zapEnc := zap.NewProductionEncoderConfig()
	zapEnc.TimeKey = "ts"
	zapEnc.MessageKey = "msg"
	zapEnc.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapEnc
}

func newZapWriter(path string, name string, suffix string, skip int) (writer *zapWriter, err error) {
	if skip == 0 {
		skip = callerSkipOffset
	}
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})
	zapEnc := getZapEncode()
	lowWriteSyncer := getRotatelogs(path, name, suffix, "info")
	highWriteSyncer := getRotatelogs(path, name, suffix, "error")
	highCore := zapcore.NewCore(zapcore.NewJSONEncoder(zapEnc), zapcore.AddSync(highWriteSyncer), highPriority)
	lowCore := zapcore.NewCore(zapcore.NewJSONEncoder(zapEnc), zapcore.AddSync(lowWriteSyncer), lowPriority)
	logger := zap.New(zapcore.NewTee(highCore, lowCore), zap.AddCaller(), zap.AddCallerSkip(skip))
	writer = &zapWriter{
		log: logger,
	}
	return
}

func (z *zapWriter) getMapAnyItems(m Any) (items []zap.Field) {
	itemCount := len(m)
	if itemCount <= 0 {
		return
	}
	isEven := itemCount%2 == 0
	if isEven {
		for index := range m {
			if index%2 != 0 {
				continue
			}
			key, ok := m[index].(string)
			if !ok {

			}
			vale := m[index+1]
			items = append(items, zap.Any(key, vale))
		}
	} else {
		items = append(items, zap.Any("info", m))
	}
	return
}

func (z *zapWriter) Debug(msg string, m ...interface{}) {
	items := z.getMapAnyItems(m)
	z.log.Debug(msg, items...)
}

func (z *zapWriter) Info(msg string, m ...interface{}) {
	items := z.getMapAnyItems(m)
	z.log.Info(msg, items...)
}

func (z *zapWriter) Warn(msg string, m ...interface{}) {
	items := z.getMapAnyItems(m)
	z.log.Warn(msg, items...)
}

func (z *zapWriter) Error(msg string, m ...interface{}) {
	items := z.getMapAnyItems(m)
	z.log.Error(msg, items...)
}

func (z *zapWriter) Fatal(msg string, m ...interface{}) {
	items := z.getMapAnyItems(m)
	z.log.Fatal(msg, items...)
}
