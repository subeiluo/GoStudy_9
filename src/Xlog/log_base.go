package Xlog

import (
	"fmt"
	"os"
	"time"
)

type LogBase struct {
	level  int
	module string
}
type Output struct {
	levelStr string
	t        string
	module   string
	filename string
	funcname string
	lineNo   int
	data     string
}

func (l *LogBase) writerlog(xfile *os.File, out *Output) { //传入输出方式,和要输出的结构
	fmt.Fprintf(xfile, "%s %s %s (%s:%s:%d) %s\n",
		out.t, out.levelStr, out.module,
		out.filename, out.funcname, out.lineNo, out.data)
}

func getTimeStr() string {
	str := time.Now().Format("2006-01-02 15:04:05.000")
	return str
}

func (l *LogBase) fomatlogger(level int, module string, xfmt string, args ...interface{}) *Output {
	levelStr := getLevelStr(level)
	t := getTimeStr()
	moduel := module
	filename, funcName, lineNo := getlineInfo(3)
	data := fmt.Sprintf(xfmt, args...) //args 需要加...展开 不然会有多余数据

	return &Output{
		t:        t,
		levelStr: levelStr,
		module:   moduel,
		filename: filename,
		funcname: funcName,
		lineNo:   lineNo,
		data:     data,
	}
}
