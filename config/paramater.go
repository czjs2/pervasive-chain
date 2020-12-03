package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"pervasive-chain/model"
)

func ReadWebCfg(path string) (*model.RuntimeConfig, error) {
	PrjConfig := &model.RuntimeConfig{}
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
