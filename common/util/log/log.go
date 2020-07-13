package LogUtil

import (
	"log"
	"fmt"
)

var (
	logCallBack func(logStr string)
)

func SetLogCallBack(logCallBackParam func(logStr string))  {
	logCallBack = logCallBackParam
}

func Println(v ...interface{})  {
	log.Print(v)
	if nil != logCallBack {
		logCallBack(fmt.Sprint(v...))
	}
}
