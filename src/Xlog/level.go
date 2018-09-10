package Xlog

import "time"

const (
	XlogLevelDebug =iota
	XlogLevelTrace
	XlogLevelInfo
	XlogLevelWarn
	XlogLevelError
	XlogLevelfatal
)
const (
	XlogTypeFile =iota   //输出到文件
	XlogTyoeConsole		 //输出到控制台
)

func getLevelStr(level int)string{
	switch level {
	case XlogLevelDebug:
		return "DEBUG"
	case XlogLevelTrace:
		return "TRACE"
	case XlogLevelInfo:
		return "INFO"
	case XlogLevelWarn:
		return "WARN"
	case XlogLevelError:
		return "ERROR"
	case XlogLevelfatal:
		return "FATAL"
	default:
		return "UNKNOW"
	}
}
func getTimeStr() string   {
	str:=time.Now().Format("2006-01-02 15:04:05.000")
	return str
}