package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/dao"
	"pervasive-chain/db"
	"pervasive-chain/form"
	"pervasive-chain/model"
)

type NodeService struct {
	dao dao.IDao
}

func (n *NodeService) ChainList() (interface{}, int, error) {
	var query []bson.M
	node := model.Node{}
	return n.dao.List(query, &node)
}

func (n *NodeService) UpdateNodeInfo(nodeForm form.HeartBeatFrom) (interface{}, error) {
	keyId := fmt.Sprintf("%s-%s", nodeForm.Type, nodeForm.Id)
	param := bson.M{
		"keyId":    keyId,
		"number":   nodeForm.Number,
		"lastTime": nodeForm.Time,
	}
	update := options.Update()
	update.SetUpsert(true)
	return n.dao.UpdateWithOption(bson.M{"keyId": keyId}, param, update)
}

func NewNodeService() INodeService {
	return &NodeService{dao: dao.NewDao(db.Node)}
}
