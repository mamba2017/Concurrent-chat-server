package myLog

import (
	"fmt"
	"os"
)

type FileLogger struct {
	level int
	logFilePath string
	logFileName string
	logFile *os.File  //os包中File类型的指针
}
//golang 中的构造函数技巧
/*
type Person {
    name string,
    age    int64,
    country string,
    ...
}

func NewPerson(name string,age int64,country sting)*Person{
      return &Person{ name: name,
}
 */


//专门用来初始化日志的文件句柄   只在文件内部使用
func  (f *FileLogger)initFileLogger()  {
	//打开日志文件
	filepath := fmt.Sprintf("%s/%s",f.logFilePath,f.logFileName)
	file ,err := os.OpenFile(filepath,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil{
		panic(fmt.Sprintf("open log: %s failed",filepath))
	}
	f.logFile = file
}
func NewFileLogger(level int,logFilePath,logFileName string) *FileLogger{
	flObj := &FileLogger{
		level: level,
		logFilePath: logFilePath,
		logFileName: logFileName,
	}
	flObj.initFileLogger()   //调用初始化方法  上面
	return flObj
}
//记录日志
func (f *FileLogger)Debug(format string, args ...interface{})  {
	//往文件里写文件
	//f.logFile.WriteString(msg)
	fmt.Fprintf(f.logFile,format,args...)
	fmt.Fprintln(f.logFile)
}
func (f *FileLogger)INFO(msg string)  {
	//往文件里写文件
	f.logFile.WriteString(msg)
}
