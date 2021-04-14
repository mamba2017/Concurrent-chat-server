package main

import "mylogger"
func main() {
	//logger := mylogger.NewFileLogger("debug","./","xxx.log")
	logger := mylogger.NewConsloleLogger("debug")
	sb := "王德法"
	logger.Debug("%s是个fuck",sb)
	logger.Info("%s是个fuck",sb)
	logger.Warn("%s是个fuck",sb)
	logger.Error("%s是个fuck",sb)
	logger.Fatal("%s是个fuck",sb)
}
