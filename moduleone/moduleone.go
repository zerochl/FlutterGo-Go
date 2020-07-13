package moduleone

import (
	"FlutterGo-Go/common/entity/request"
	"encoding/json"
)

func AAddB(requestJson string) (int, error) {
	helloWorld := &request.HelloWorld{}
	err := json.Unmarshal([]byte(requestJson), helloWorld)
	if err != nil {
		return 0, err
	}
	return helloWorld.A + helloWorld.B, nil
}
