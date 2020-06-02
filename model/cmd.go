package model

// 下发执行命令
type Cmd struct {
	Key    string
	Params map[string]interface{}
}
