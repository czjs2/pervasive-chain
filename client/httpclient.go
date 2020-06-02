package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"pervasive-chain/form"
	"strings"
)

/**
向lotus节点获取数据的疯转 client
*/

var client *http.Client

func NewHttpClient() *http.Client {
	if client == nil {
		client = http.DefaultClient
	}
	return client
}

func NewRequest(host, data, token string) (string, error) {
	request, err := http.NewRequest("POST", host, strings.NewReader(data))
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)
	res, err := NewHttpClient().Do(request)
	if err != nil {
		return "", err
	}
	if res.StatusCode < 200 || res.StatusCode > 300 {
		return "", errors.New("response fail ")
	}
	if res == nil {
		return "", errors.New("response is nil")
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(fmt.Sprintf("------------------- %s ---------------------",request.URL))
	fmt.Println(data)
	fmt.Println(string(bytes))
	fmt.Println("----------------------end-----------------------")
	return string(bytes), nil
}

/**
HeartBeat
*/

func HeartBeat(host, path, token string, heartForm form.HeartBeatFrom) (string, error) {
	bytes, _ := json.Marshal(heartForm)
	return NewRequest(fmt.Sprintf("%s%s", host, path), string(bytes), token)
}

/**
ReportBlock
*/

func ReportBlock(host, path, token string, blockForm form.ReportBlockForm) (string, error) {
	bytes, _ := json.Marshal(blockForm)
	return NewRequest(fmt.Sprintf("%s%s", host, path), string(bytes), token)
}

func ReportFlow(host, path, token string, flowForm form.ReportFlowForm) (string, error) {
	bytes, _ := json.Marshal(flowForm)
	return NewRequest(fmt.Sprintf("%s%s", host, path), string(bytes), token)
}
