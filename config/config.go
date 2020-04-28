package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type WebConfig struct {
	LogPath       string
	HTTPListen    string
	Debug         bool // 是否是debug
	MongodbUrl    string
	DevMongodbUrl string
	WebRoot       string // 静态资源路径
	HtmlTemplate  string
}

type ChinConfig struct {
	LogPath       string
	HTTPListen    string
	Debug         bool // 是否是debug
	MongodbUrl    string
	DevMongodbUrl string
}

var PrjConfig *WebConfig

func ReadCfg(path string) (*WebConfig, error) {
	PrjConfig := &WebConfig{}
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(d))
	err = json.Unmarshal(d, &PrjConfig)
	if err != nil {
		return PrjConfig, err
	}
	return PrjConfig, nil
}
