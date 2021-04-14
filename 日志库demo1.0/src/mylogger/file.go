package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志

//文件日志结构体
type FileLogger struct {
	level Level
	fileName string
	filePath string
	file *os.File   //打开文件句柄
	errFile *os.File
}
//文件日志结构体的构造函数
func NewFileLogger(levelStr,fileName,filePath string)(*FileLogger){
	logLevel := parseLogLevel(levelStr)
	fl:= &FileLogger{
		level: logLevel,
		fileName: fileName,
		filePath: filePath,
	}
	fl.initFile()   //根据文件路径和文件名打开日志文件，把文件句柄赋值给结构体对应的字段
	return fl
}
//将指定的日志文件打开，赋值给结构体
func (f *FileLogger)initFile()  {
	logName := path.Join(f.filePath,f.fileName)
	//打开文件
	fileObj,err := os.OpenFile(logName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0664)
	if err != nil{
		panic(fmt.Errorf("Open %s file failed,%v",logName,err))
	}
	f.file = fileObj

//打开错误日志文件
	errLogName := fmt.Sprintf(".err",logName)
	errFileObj,err := os.OpenFile(errLogName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0664)
	if err != nil{
		panic(fmt.Errorf("Open %s errfile failed,%v",errLogName,err))
	}
	f.errFile = errFileObj
}
func (f *FileLogger)log(level Level,format string,args... interface{})  {
	if f.level > level{
		return
	}
	loglevelstr := getLevelStr(level)
	msg := fmt.Sprintf(format,args...)   //得到用户要记录的日志
	//日志格式：【时间】【文件：行号】【函数名】【日志级别】 日志信息
	now := time.Now().Format("2006-01-02 15:04:05")
	fileName,funcName ,line :=  getCallerInfo(3)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s",now,fileName,line,funcName,loglevelstr,msg)
	fmt.Fprintln(f.file,logMsg) //将msg字符串写入f.file文件
	//如果是error或者fatal级别的日志还要记录到f.errFile
	if level >= ErrorLevel{
		fmt.Fprintln(f.errFile,logMsg)
	}
}


//debug方法
func (f *FileLogger)Debug(format string,args... interface{}){
	////f.file.Write()
	//if f.level > DebugLevel{
	//	return
	//}
	//msg := fmt.Sprintf(format,args...)   //得到用户要记录的日志
	////日志格式：【时间】【文件：行号】【函数名】【日志级别】 日志信息
	//now := time.Now().Format("2006-01-02 15:04:05")
	//fileName,funcName ,line :=  getCallerInfo(2)
	//logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s\n",now,fileName,line,funcName,"debug",msg)
	//fmt.Fprintf(f.file,logMsg) //将msg字符串写入f.file文件
	////fmt.Errorf()
	////fmt.Sprintf()
	f.log(DebugLevel,format,args...)
}
//info方法
func (f *FileLogger)Info(format string,args... interface{}){
	////f.file.Write()
	//if f.level > InfoLevel{
	//	return
	//}
	//msg := fmt.Sprintf(format,args...)   //得到用户要记录的日志
	////日志格式：【时间】【文件：行号】【函数名】【日志级别】 日志信息
	//now := time.Now().Format("2006-01-02 15:04:05")
	//fileName,funcName ,line :=  getCallerInfo(2)
	//logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s\n",now,fileName,line,funcName,"info",msg)
	//fmt.Fprintf(f.file,logMsg) //将msg字符串写入f.file文件
	f.log(InfoLevel,format,args...)
}
//warn方法
func (f *FileLogger)Warn(format string,args... interface{}){
	////f.file.Write()
	//if f.level > InfoLevel{
	//	return
	//}
	//msg := fmt.Sprintf(format,args...)   //得到用户要记录的日志
	////日志格式：【时间】【文件：行号】【函数名】【日志级别】 日志信息
	//now := time.Now().Format("2006-01-02 15:04:05")
	//fileName,funcName ,line :=  getCallerInfo(2)
	//logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s\n",now,fileName,line,funcName,"info",msg)
	//fmt.Fprintf(f.file,logMsg) //将msg字符串写入f.file文件
	f.log(WarningLevel,format,args...)
}
//error 方法
func (f *FileLogger)Error(format string,args... interface{}){
	////f.file.Write()
	//if f.level > InfoLevel{
	//	return
	//}
	//msg := fmt.Sprintf(format,args...)   //得到用户要记录的日志
	////日志格式：【时间】【文件：行号】【函数名】【日志级别】 日志信息
	//now := time.Now().Format("2006-01-02 15:04:05")
	//fileName,funcName ,line :=  getCallerInfo(2)
	//logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s\n",now,fileName,line,funcName,"info",msg)
	//fmt.Fprintf(f.file,logMsg) //将msg字符串写入f.file文件
	f.log(ErrorLevel,format,args...)
}
//fatal 方法
func (f *FileLogger)Fatal(format string,args... interface{}){
	////f.file.Write()
	//if f.level > InfoLevel{
	//	return
	//}
	//msg := fmt.Sprintf(format,args...)   //得到用户要记录的日志
	////日志格式：【时间】【文件：行号】【函数名】【日志级别】 日志信息
	//now := time.Now().Format("2006-01-02 15:04:05")
	//fileName,funcName ,line :=  getCallerInfo(2)
	//logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s\n",now,fileName,line,funcName,"info",msg)
	//fmt.Fprintf(f.file,logMsg) //将msg字符串写入f.file文件
	f.log(FatalLevel,format,args...)
}
