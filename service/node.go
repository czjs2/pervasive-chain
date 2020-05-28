package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"pervasive-chain/dao"
	"pervasive-chain/db"
	"pervasive-chain/form"
)

type NodeService struct {
	dao dao.IDao
}

/**
keyId:String, //节点id 类型+key [b|r|s]-[Id]
	number:String,//链编号
	lastTime:Date,//最近一次上报心跳时间
	cmd:{key:String,params:Object},//命令 :{执行码，参数}
	cmdTime:Date, //命令产生时间
*/
func (n *NodeService) UpdateNodeInfo(nodeForm form.HeartBeatFrom) (interface{}, error) {
	param := bson.M{
		"keyId":fmt.Sprintf("%s-%s",nodeForm.Type,nodeForm.Id),
		"number":nodeForm.Number,
		"lastTime":nodeForm.Time,
	}
	return n.dao.Add(param)
}

func NewNodeService() INodeService {
	return &NodeService{dao: dao.NewDao(db.Node)}
}
