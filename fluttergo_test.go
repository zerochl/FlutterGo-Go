package fluttergo

import (
	"testing"
	"log"
)

type LogCallBackTest struct {

}

func (lcbt *LogCallBackTest) OnLog(logStr string)  {
	log.Print("test log:", logStr)
}

func TestHelloWorld(t *testing.T) {
	logUtil := &LogCallBackTest{}
	Init(logUtil)
	log.Println(HelloWorld("{\"a\":1,\"b\":2}"))
}
