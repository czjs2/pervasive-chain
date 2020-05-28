package service

import "pervasive-chain/form"

type INodeService interface {
	UpdateNodeInfo(nodeForm form.HeartBeatFrom) (interface{}, error)
}
