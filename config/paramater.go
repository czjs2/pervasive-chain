package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Param map[string]interface{}

func ReadWebCfg(path string) (*RuntimeConfig, error) {
	PrjConfig := &RuntimeConfig{}
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(d))
	err = json.Unmarshal(d, PrjConfig)
	if err != nil {
		return PrjConfig, err
	}
	return PrjConfig, nil
}



type RuntimeConfig struct {
	Debug         bool // 是否是debug
	LogPath       string
	HTTPListen    string
	MongodbUrl    string
	DevMongodbUrl string
}
