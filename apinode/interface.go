package apinode

import "pervasive-chain/form"

type INodeService interface {
	UpdateNodeInfo(heartFrom form.HeartBeatFrom) (interface{}, error)
}
