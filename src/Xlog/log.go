package Xlog

type Xlog interface {
	LogDebug(xfmt string, args...interface{})
	LogTrace(xfmt string, args...interface{})
	LogInfo(xfmt string, args...interface{})
	LogWarn(xfmt string, args...interface{})
	LogError(xfmt string, args...interface{})
	LogFatal(xfmt string, args...interface{})

	SetLevel(level int) 
}

func NewXlog(logType,level int,filename string,module string)Xlog  {
	var logger Xlog
	switch (logType) {
	case XlogTypeFile:
		logger = NewFile(level,filename,module)

	case XlogTyoeConsole:
		logger= NewXconsole(level,module)
	default:
		logger = NewFile(level,filename,module)
	}
	return logger
}