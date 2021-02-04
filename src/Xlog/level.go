package Xlog

const (
	ClogLevelDebug = iota  //0
	ClogLevelTrace
	ClogLevelInfo
	ClogLevelWarn
	ClogLevelError
	ClogLevelFatal
)
const (
	ClogTypeFile    = iota //输出到文件
	ClogTypeConsole        //输出到控制台
)

func getLevelStr(level int) string {
	switch level {
	case ClogLevelDebug:
		return "DEBUG"
	case ClogLevelTrace:
		return "TRACE"
	case ClogLevelInfo:
		return "INFO"
	case ClogLevelWarn:
		return "WARN"
	case ClogLevelError:
		return "ERROR"
	case ClogLevelFatal:
		return "FATAL"
	default:
		return "UNKNOW"
	}
}
