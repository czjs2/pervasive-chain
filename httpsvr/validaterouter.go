package httpsvr

import (
	"bytes"
	"fmt"
	"pervasive-chain/config"
	"pervasive-chain/service"
	"sync"
)

type router map[string]func(req string) (service.IFormValidateInterface, error)

var validateManager = NewValidateRouter()

type ValidateRouter struct {
	router router
	sync.Mutex
}

func (wsd *ValidateRouter) RouterInfo() string {
	var buff bytes.Buffer
	for k, v := range wsd.router {
		buff.WriteString(fmt.Sprintf("%v %v ", k, v))
	}
	return buff.String()
}

func (wsd *ValidateRouter) Exists(path string) bool {
	wsd.Lock()
	defer wsd.Unlock()
	_, ok := wsd.router[path]
	return ok
}

func (wsd *ValidateRouter) Register(path string, fn func(req string) (service.IFormValidateInterface, error)) {
	path = fmt.Sprintf("%v%v", config.ApiVersion, path)
	ok := wsd.Exists(path)
	if ok {
		panic(fmt.Sprintf("validate router info have exists %v \n", path))
	}
	wsd.Lock()
	defer wsd.Unlock()
	wsd.router[path] = fn
}

func (wsd *ValidateRouter) Execute(path string, req string) error {
	ok := wsd.Exists(path)
	if !ok {
		return fmt.Errorf("validate router path is not exitsts %v \n", path)
	}
	wsHandler := wsd.router[path]
	objForm, err := wsHandler(req)
	if err != nil {
		return err
	}
	ok, err = objForm.Valid()
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("param error \n")
	}
	return nil

}

func NewValidateRouter() ValidateRouter {
	return ValidateRouter{router: router{}}
}
