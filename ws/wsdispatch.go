package ws

import (
	"bytes"
	"fmt"
	"pervasive-chain/service"
	"sync"
)

var WsRouterManaer = NewWsDispatch()

type WsRouter map[string]func(c *WsContext)

type ValidateRouter map[string]func(c *WsContext) (service.IFormValidateInterface, error)

type WsDispatch struct {
	router         WsRouter
	validateRouter ValidateRouter
	sync.Mutex
}

func (wsd *WsDispatch) RouterInfo() string {
	var buff bytes.Buffer
	for k, v := range wsd.router {
		buff.WriteString(fmt.Sprintf("%v %v ", k, v))
	}
	return buff.String()
}

func (wsd *WsDispatch) Exists(path string) bool {
	wsd.Lock()
	defer wsd.Unlock()
	_, ok := wsd.router[path]
	return ok
}

func (wsd *WsDispatch) Register(path string, fn func(c *WsContext), v func(c *WsContext) (service.IFormValidateInterface, error)) {
	ok := wsd.Exists(path)
	if ok {
		panic(fmt.Sprintf("websocket router info have exists %v \n", path))
	}
	wsd.Lock()
	defer wsd.Unlock()
	wsd.router[path] = fn
	if v != nil {
		wsd.validateRouter[path] = v
	}
}

func (wsd *WsDispatch) Execute(path string, c *WsContext) error {
	ok := wsd.Exists(path)
	if !ok {
		return fmt.Errorf("websocket router path is not exitsts %v \n", path)
	}
	v, ok := wsd.validateRouter[path]
	if ok {
		fromValid, err := v(c)
		if err!=nil{
			return err
		}
		valid, err := fromValid.Valid()
		if err!=nil{
			return err
		}
		if !valid{
			return fmt.Errorf("websocket params error %v \n",path)
		}

	}

	wsHandler := wsd.router[path]
	wsHandler(c)
	return nil
}

func NewWsDispatch() WsDispatch {
	return WsDispatch{
		router:         WsRouter{},
		validateRouter: ValidateRouter{},
	}
}
