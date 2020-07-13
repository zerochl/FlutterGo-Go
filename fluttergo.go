package fluttergo

import (
	"log"
	"FlutterGo-Go/common/entity/response"
	"FlutterGo-Go/common/constants"
	"FlutterGo-Go/common/util/log"
	"FlutterGo-Go/moduleone"
)

func HelloWorld(requestJson string) (result string) {
	LogUtil.Println("in HelloWorld")
	defer catchError(&result)
	sum, err := moduleone.AAddB(requestJson)
	if err != nil {
		return response.NewBaseResponse(constants.NormalErrorCode, constants.NormalErrorMsg + err.Error(), nil).ToJson()
	}
	return response.NewBaseResponse(constants.Success, "", sum).ToJson()
}

func catchError(result *string)  {
	if err := recover(); err != nil {
		log.Println("error:", err.(error).Error())
		*result = response.NewBaseResponse(constants.NormalErrorCode, constants.NormalErrorMsg + err.(error).Error(), nil).ToJson()
	}
}