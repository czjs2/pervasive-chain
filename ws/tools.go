package ws

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"github.com/tidwall/gjson"
)

func GetJsonValue(src string, filed string) string {
	value := gjson.Get(src, filed)
	return value.String()
}

func GetJsonArray(src string, filed string) *[]gjson.Result {
	results := gjson.GetMany(src, filed)
	return &results
}

func Unmarshal(src string, obj interface{}) error {
	return json.Unmarshal([]byte(src), obj)
}

// 获取唯一的Id
func GetUUID() string {
	id := uuid.NewV4()
	return id.String()
}
