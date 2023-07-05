package llog

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"time"
)

func getRotatelogs(path string, name string, suffix string, logType string) io.Writer {
	// 每1小时(整点)分割一次日志
	if path[len(path)-1] != '/' {
		path += "/"
	}
	filename := path + name + "/%Y%m%d/" + name + "_" + logType + "_%H.log"
	if suffix != "" {
		filename = path + name + "/%Y%m%d/" + name + "_" + suffix + "_" + logType + "_%H.log"
	}

	hook, err := rotatelogs.New(
		filename,
		//rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
