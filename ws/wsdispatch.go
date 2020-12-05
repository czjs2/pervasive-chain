package ws

import (
	"bytes"
	"fmt"
	"sync"
)

var WsRouterManaer = NewWsDispatch()

type WsRouter map[string]func(c *WsContext)

type WsDispatch struct {
	router WsRouter
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

func (wsd *WsDispatch) Register(path string, fn func(c *WsContext)) {
	ok := wsd.Exists(path)
	if ok {
		panic(fmt.Sprintf("websocket router info have exists %v \n", path))
	}
	wsd.Lock()
	defer wsd.Unlock()
	wsd.router[path] = fn
}

func (wsd *WsDispatch) Execute(path string, c *WsContext) error {
	ok := wsd.Exists(path)
	if !ok {
		return fmt.Errorf("websocket router path is not exitsts %v \n", path)
	}
	wsHandler := wsd.router[path]
	wsHandler(c)
	return nil
}

func NewWsDispatch() WsDispatch {
	return WsDispatch{router: WsRouter{}}
}
