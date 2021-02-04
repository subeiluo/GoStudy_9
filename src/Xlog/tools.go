package Xlog

import (
	"runtime"
)

//获取文件的文件名 函数名 和调用的行数

func getlineInfo(skip int) (filename, funcname string, lineNO int) {
	//skip的值根据代码的层数深度来确定 如main 是1 xlog.GetlineInfo是2 向下判断
	pc, file, line, ok := runtime.Caller(skip)  //skip 0
	if ok {
		fun := runtime.FuncForPC(pc)
		funcname = fun.Name()
	}
	filename = file
	lineNO = line
	return
}