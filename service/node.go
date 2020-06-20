package service

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/dao"
	"pervasive-chain/db"
	"pervasive-chain/form"
	"pervasive-chain/model"
	"time"
)

type NodeService struct {
	dao dao.IDao
}

func (n *NodeService) UpdateOnLineNodeCmd(cmd model.PyCmd) (interface{}, error) {
	query := bson.M{

	}
	param := bson.M{
		"cmd":     cmd,
		"cmdTime": time.Now(),
	}
	return n.dao.UpdateMany(query, param)
}

func (n *NodeService) OnLineList() (interface{}, int, error) {
	// todo 在
	query := []bson.M{
		bson.M{"$match": bson.M{}},
	}
	return n.dao.List(query)
}

func (n *NodeService) ChainList() (interface{}, int, error) {
	var query []bson.M
	return n.dao.List(query)
}

func (n *NodeService) FindAndUpdate(nodeForm form.HeartBeatFrom) (*model.Node, error) {
	keyId := fmt.Sprintf("%s-%s", nodeForm.Type, nodeForm.Id)
	if nodeForm.Time == 0 {
		return nil, errors.New("time is zero ")
	}
	query := bson.M{"keyId": keyId}
	param := bson.M{
		"type":     nodeForm.Type,
		"keyId":    keyId,
		"number":   nodeForm.Number,
		"lastTime": millisecondToTime(nodeForm.Time).Local(),
		//"cmd":      nil,
		//"cmdTime":  nil,
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
