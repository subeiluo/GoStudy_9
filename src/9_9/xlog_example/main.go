package main

import (
	"Xlog"
	"flag"
	"fmt"
)

func logic(logger Xlog.Xlog) {
	logger.LogDebug("debug log")
	logger.LogTrace("trace log")
	logger.LogInfo("info log")
	logger.LogWarn("warn log")
	logger.LogError("error log")
	logger.LogFatal("fatal log")
	logger.LogDebug("debug log")
}

func main() {
	//filename,funcName,lineNo :=Xlog.GetlineInfo(1)   //skip的值根据代码的层数深度来确定 如main 是1 xlog.GetlineInfo是2 向下判断
	//fmt.Printf("filename:%s funcName:%s line:%d\n",filename,funcName,lineNo)
	var logTypeStr string
	flag.StringVar(&logTypeStr, "type", "file", "please input logger type")
	flag.Parse()

	var logType int
	if logTypeStr == "file" {
		logType = Xlog.ClogTypeFile
	} else {
		logType = Xlog.ClogTypeConsole
	}
	//logType 传入输出方式 ,日志级别, 文件位置, 模块信息
	logger := Xlog.NewXlog(logType, Xlog.ClogLevelDebug,
		"/Users/subeiluo/Downloads/logtest.log",
		"xlog_example")
	err := logger.Init()
	if err != nil {
		fmt.Printf("logger init failed\n")
		return
	}
	logic(logger)
}
