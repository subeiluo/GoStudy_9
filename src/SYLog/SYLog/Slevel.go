package SYLog

//定义常量日志级别 使用iota 自增赋值
const (
	DebugLevelLog = iota
	TraceLevelLog
	InfoLevelLog
	WarnLevelLog
	ErrorLevelLog
	FatalLevelLog
)

func getLoglevel(level int)string{
	switch level {
	case DebugLevelLog:
		return "DEBUG"
	case TraceLevelLog:
		return "TRACE"
	case InfoLevelLog:
		return  "INFO"
	case WarnLevelLog:
		return  "WARN"
	case ErrorLevelLog:
		return  "ERROR"
	case FatalLevelLog:
		return  "FATAL"
	default:
		return "UNKNOW"
	}
}
