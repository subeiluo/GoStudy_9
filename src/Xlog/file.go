package Xlog

import (
	"os"
	"sync"
)

type XFile struct {
	*LogBase
	filename string
	file     *os.File

	logChan chan *Output
	wg *sync.WaitGroup
}

func NewFile(level int, filename string, module string) Xlog {
	logger := &XFile{
		filename: filename,
	}
	logger.LogBase = &LogBase{
		level:  level,
		module: module,
	}
	logger.logChan =make(chan *Output, 10000)
	go logger.synclog()
	return logger
}
func(c *XFile)synclog(){
	for data :=range c.logChan{
		c.writerlog(c.file,data)
	}
	c.wg.Done()
}
func (c *XFile) Init() (err error) {
	c.file, err = os.OpenFile(c.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	//文件名,追加方式写入,创建文件,写入权限0755 rwxr-xr-xr
	if err != nil {
		return
	}
	return
}

func (c *XFile) LogDebug(xfmt string, args ...interface{}) {
	if c.level > ClogLevelDebug {
		return
	}
	logData := c.fomatlogger(c.level, c.module, xfmt, args...)
	select {
		case c.logChan <- logData:
	default:

	}
	//传入文件句柄
	c.writerlog(c.file, logData)
}
func (c *XFile) LogTrace(xfmt string, args ...interface{}) {
	if c.level > ClogLevelTrace {
		return
	}
	logData := c.fomatlogger(c.level, c.module, xfmt, args...)
	logData.levelStr = "TRACE"
	c.writerlog(c.file, logData)
}
func (c *XFile) LogInfo(xfmt string, args ...interface{}) {
	if c.level > ClogLevelInfo {
		return
	}
	logData := c.fomatlogger(c.level, c.module, xfmt, args...)
	logData.levelStr = "INFO"
	c.writerlog(c.file, logData)
}
func (c *XFile) LogWarn(xfmt string, args ...interface{}) {
	if c.level > ClogLevelWarn {
		return
	}
	logData := c.fomatlogger(c.level, c.module, xfmt, args...)
	logData.levelStr = "WARN"
	c.writerlog(c.file, logData)
}
func (c *XFile) LogError(xfmt string, args ...interface{}) {
	if c.level > ClogLevelError {
		return
	}
	logData := c.fomatlogger(c.level, c.module, xfmt, args...)
	logData.levelStr = "ERROR"
	c.writerlog(c.file, logData)
}
func (c *XFile) LogFatal(xfmt string, args ...interface{}) {
	if c.level > ClogLevelFatal {
		return
	}
	logData := c.fomatlogger(c.level, c.module, xfmt, args...)
	logData.levelStr = "FATAL"
	c.writerlog(c.file, logData)
}
func (c *XFile) SetLevel(level int) {
	c.level = level
}
func (c *XFile) Close() {
	if c.file != nil {
		c.file.Close()
	}

}
