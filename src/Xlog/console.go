package Xlog

import (
	"os"
)

type Xconsole struct {
	*LogBase
}

func NewXconsole(level int, module string) Xlog {
	logger := &Xconsole{}
	logger.LogBase = &LogBase{
		level:  level,
		module: module,
	}
	return logger
}
func (c *Xconsole) Init()(err error) {
	return nil
}
func (c *Xconsole) LogDebug(xfmt string, args ...interface{}) {
	logData := c.fomatlogger(ClogLevelDebug, c.module, xfmt, args...)
	logData.levelStr = "DEBUG"
	//传入控制台输出句柄
	c.writerlog(os.Stdout, logData)
}
func (c *Xconsole) LogTrace(xfmt string, args ...interface{}) {
	logData := c.fomatlogger(ClogLevelTrace, c.module, xfmt, args...)
	logData.levelStr = "TRACE"
	c.writerlog(os.Stdout, logData)
}
func (c *Xconsole) LogInfo(xfmt string, args ...interface{}) {
	logData := c.fomatlogger(ClogLevelInfo, c.module, xfmt, args...)
	logData.levelStr = "INFO"
	c.writerlog(os.Stdout, logData)
}
func (c *Xconsole) LogWarn(xfmt string, args ...interface{}) {
	logData := c.fomatlogger(ClogLevelWarn, c.module, xfmt, args...)
	logData.levelStr = "WARN"
	c.writerlog(os.Stdout, logData)
}
func (c *Xconsole) LogError(xfmt string, args ...interface{}) {
	logData := c.fomatlogger(ClogLevelError, c.module, xfmt, args...)
	logData.levelStr = "ERROR"
	c.writerlog(os.Stdout, logData)
}
func (c *Xconsole) LogFatal(xfmt string, args ...interface{}) {
	logData := c.fomatlogger(ClogLevelFatal, c.module, xfmt, args...)
	logData.levelStr = "FATAL"
	c.writerlog(os.Stdout, logData)
}
func (c *Xconsole) SetLevel(level int) {
	c.level = level
}
func (c *Xconsole) Close() {
}
