package Xlog

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type XFile struct {
	*LogBase
	filename string
	file     *os.File   //指针类型  os包中file类型指针

	logChan chan *Output
	wg *sync.WaitGroup
	nowtime int
}

func NewFile(level int, filename string, module string) Xlog {
	logger := &XFile{
		filename: filename,
	}
	logger.LogBase = &LogBase{
		level:  level,
		module: module,
	}
	logger.nowtime = time.Now().Hour()
	logger.logChan =make(chan *Output, 10000)
	go logger.synclog()
	return logger
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
	// c.level == c.LogBase.level
	if c.LogBase.level > ClogLevelDebug {
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

func(c *XFile)synclog(){
	for data :=range c.logChan{

		c.writerlog(c.file,data)
	}
	c.wg.Done()
}

func (c *XFile)timesplitlog(){
	nowhour :=time.Now()
	//获取当前时间 并且和日志文件初始化的时间相比较
	//相同则忽略
	if nowhour.Hour() ==c.nowtime{
		return
	}
	//将当前内存中的数据进行落地后 关闭文件
	c.file.Sync()
	//关闭文件前将当前时间进行更新
	c.nowtime =nowhour.Hour()
	c.file.Close()
	newfilename :=fmt.Sprintf("%s-%04d-%02d-%02d-%02d",c.filename,nowhour.Year(),nowhour.Month(),nowhour.Day(),nowhour.Hour())
	os.Rename(c.filename,newfilename)
	c.Init()
}

const (
	bit =1024
	kb  =bit*1000
	mb  = kb *1000
)

/*func(c *XFile)sizesplitlog(){
	sz,err :=os.Stat(c.filename)
	if err!=nil{
		return
	}
	if sz.Size() >= 50*mb{
		nowhour :=time.Now()
		c.file.Sync()
		//关闭文件前将当前时间进行更新
		c.file.Close()
		newfilename :=fmt.Sprintf("%s-%04d-%02d-%02d-%02d",c.filename,nowhour.Year(),nowhour.Month(),nowhour.Day(),nowhour.Hour())
		os.Rename(c.filename,newfilename)
		c.Init()
	}

}

 */