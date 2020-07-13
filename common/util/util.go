package util

import "encoding/json"

func ToJson(obj interface{}) string {
	data, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(data)
}