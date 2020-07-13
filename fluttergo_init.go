package fluttergo

import (
	"FlutterGo-Go/common/util/log"
)

type LogCallBack interface {
	OnLog(logStr string)
}

func Init(logCallBack LogCallBack)  {
	LogUtil.SetLogCallBack(logCallBack.OnLog)
}
