package SYLog

import "time"

type LogBase struct {
	level int
	module string
}

type OutPut struct {
	time string
	levelstr string
	fileName string
	funcName string
	line int
	data string
}

func getTimeStr()string  {
	str:= time.Now().Format("[2006-01-02 15:04:05]")
	return str
}

func (l *LogBase)fmtLog(level int, module string,xfmt string,arge...interface{}) *OutPut {
	levelstr:=getLoglevel(level)
	ts :=getTimeStr()
	mod :=module

}