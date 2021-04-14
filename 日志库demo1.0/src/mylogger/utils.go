package mylogger

import (
	"path"
	"runtime"
)

//存放一些公用工具的函数
func getCallerInfo(skip int)(fileName,funcName string,line int){
	pc , fileName , line , ok  := runtime.Caller(skip)
	if !ok{
		return
	}
	//从filename中剥离文件名
	fileName = path.Base(fileName)
	//从pc拿到函数名
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	return
}
