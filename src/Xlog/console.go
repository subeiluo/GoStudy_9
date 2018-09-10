package Xlog

import "fmt"

type Xconsole struct {
	level int
	module string
}

func NewXconsole(level int,module string)Xlog  {
	logger :=&Xconsole{
		level:level,
		module:module,
	}
	return logger
}

func (c *Xconsole)LogDebug(xfmt string, args ...interface{}){
	fmt.Printf("log debug of console\n")
}
func (c *Xconsole)LogTrace(xfmt string, args ...interface{}){
	fmt.Printf("log trace of console\n")
}
func (c *Xconsole)LogInfo(xfmt string, args ...interface{}){
	fmt.Printf("log info of console\n")
}
func (c *Xconsole)LogWarn(xfmt string, args ...interface{}){
	fmt.Printf("log warn of console\n")
}
func (c *Xconsole)LogError(xfmt string, args ...interface{}){
	fmt.Printf("log error of console\n")
}
func (c *Xconsole)LogFatal(xfmt string, args ...interface{}){
	fmt.Printf("log fatal of console\n")
}
func (c *Xconsole)SetLevel(level int){

}