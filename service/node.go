package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/dao"
	"pervasive-chain/db"
	"pervasive-chain/form"
	"pervasive-chain/model"
	"pervasive-chain/utils"
)

type NodeService struct {
	dao dao.IDao
}

func (n *NodeService) UpdateOnLineNodeCmd(cmd model.PyCmd) (interface{}, error) {
	query := bson.M{

	}
	param := bson.M{
		"cmd":     utils.MapToStr(cmd.Key),
		"cmdTime": utils.GetNowTime(),
	}
	return n.dao.UpdateMany(query, param)
}

func (n *NodeService) OnLineList() (interface{}, int, error) {
	// todo 在
	query := []bson.M{
		bson.M{"$match": bson.M{}},
	}
	node := model.Node{}
	return n.dao.List(query, &node)
}

func (n *NodeService) ChainList() (interface{}, int, error) {
	var query []bson.M
	node := model.Node{}
	return n.dao.List(query, &node)
}

func (n *NodeService) FindAndUpdate(nodeForm form.HeartBeatFrom) (*model.Node, error) {
	keyId := fmt.Sprintf("%s-%s", nodeForm.Type, nodeForm.Id)
	query := bson.M{"keyId": keyId}
	param := bson.M{
		"type":     nodeForm.Type,
		"keyId":    keyId,
		"number":   nodeForm.Number,
		"lastTime": nodeForm.Time,
		"cmd":      nil,
		"cmdTime":  "",
	}
	update := options.FindOneAndUpdate()
	update.SetUpsert(true)
	node := model.Node{}
	_, err := n.dao.FindAndUpdate(query, param, update, &node)
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func NewNodeService() INodeService {
	return &NodeService{dao: dao.NewDao(db.Node)}
}
