package Xlog

import "runtime"

//获取文件的文件名 函数名 和调用的行数

func getlineInfo(skip int) (filename, funcname string, lineNO int) {
	pc, file, line, ok := runtime.Caller(skip) //skip 0
	if ok {
		fun := runtime.FuncForPC(pc)
		funcname = fun.Name()
	}
	filename = file
	lineNO = line
	return
}