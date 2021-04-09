package main

import (
	"myLog"
)

func main() {
	fl := myLog.NewFileLogger(myLog.INFO,"./","test.log")
	fl.Debug("测试日志！！")
	usrId := 10
	fl.Debug("usrId:%d尝试五次登陆",usrId)
}
