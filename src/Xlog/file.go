package Xlog

import (
	"fmt"
)

type XFile struct {
	level int
	filename string
	module string
}

func NewFile(level int,filename string,module string)Xlog{
	logger :=&XFile{
		level:level,
		filename:filename,
		module:module,
	}
	return logger
}

func (c *XFile)LogDebug(xfmt string, args ...interface{}) {
		levelStr :=getLevelStr(XlogLevelDebug)
		t := getTimeStr()
		moduel:=c.module
		filename,funcName,lineNo :=getlineInfo(2)
		data :=fmt.Sprintf(xfmt,args...)      					//args 需要加...展开 不然会有多余数据
		fmt.Printf("%s %v %v (%s:%s:%d) %s\n",t,levelStr,moduel,filename,funcName,lineNo,data)
}
func (c *XFile)LogTrace(xfmt string, args ...interface{}) {
	//levelStr :=getLevelStr(XlogLevelTrace)
	t := getTimeStr()
	fmt.Printf("%s log trace of file\n",t)
}
func (c *XFile)LogInfo(xfmt string, args ...interface{}) {
	fmt.Printf("log info of file\n")
}
func (c *XFile)LogWarn(xfmt string, args ...interface{}) {
	fmt.Printf("log warn of file\n")
}
func (c *XFile)LogError(xfmt string, args ...interface{}) {
	fmt.Printf("log error of file\n")
}
func (c *XFile)LogFatal(xfmt string, args ...interface{}) {
	fmt.Printf("log fatal of file\n")
}
func (c *XFile)SetLevel(level int) {

}