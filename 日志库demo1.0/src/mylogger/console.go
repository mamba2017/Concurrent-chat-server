package mylogger

import (
	"fmt"
	"os"
	"time"
)

//往终端打印日志

//终端日志结构体
type ConsloleLogger struct {
	level Level
	//file *os.File
	//errFile os.File

}
//文件日志结构体的构造函数
func NewConsloleLogger(levelStr string) *ConsloleLogger{
	logLevel := parseLogLevel(levelStr)
	cl:= &ConsloleLogger{
		level: logLevel,
		//file : os.Stdout,
		//errFile : os.Stderr,
	}
	return cl
}
func (c *ConsloleLogger)log(level Level,format string,args... interface{})  {
	if c.level > level{ //如果传进来的日志等级小于默认打印日志的等级 则不打印日志
		return
	}
	loglevelstr := getLevelStr(level)
	msg := fmt.Sprintf(format,args...)   //得到用户要记录的日志
	//日志格式：【时间】【文件：行号】【函数名】【日志级别】 日志信息
	now := time.Now().Format("2006-01-02 15:04:05")
	fileName,funcName ,line :=  getCallerInfo(3)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s",now,fileName,line,funcName,loglevelstr,msg)
	fmt.Fprintln(os.Stdout,logMsg) //将msg字符串写入f.file文件
	//如果是error或者fatal级别的日志还要记录到f.errFile

}


//debug方法
func (c *ConsloleLogger)Debug(format string,args... interface{}){

	c.log(DebugLevel,format,args...)
}
//info方法
func (c *ConsloleLogger)Info(format string,args... interface{}){

	c.log(InfoLevel,format,args...)
}
//warn方法
func (c *ConsloleLogger)Warn(format string,args... interface{}){

	c.log(WarningLevel,format,args...)
}
//error 方法
func (c *ConsloleLogger)Error(format string,args... interface{}){

	c.log(ErrorLevel,format,args...)
}
//fatal 方法
func (c *ConsloleLogger)Fatal(format string,args... interface{}){

	c.log(FatalLevel,format,args...)
}
