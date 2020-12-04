package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/dao"
	"pervasive-chain/model"
	"pervasive-chain/mongodb"
)

type NodeDao struct {
	dao mongodb.IDao
}

func (n *NodeDao) FindOne(nodeId string) (*model.Node, error) {
	obj := &model.Node{}
	_, err := n.dao.FindOne(context.TODO(), bson.M{"keyId": nodeId}, obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (n *NodeDao) UpdateLatestTime(nodeId, latestTime string) (interface{}, error) {
	update := options.Update()
	update.SetUpsert(true)
	return n.dao.UpdateWithOption(context.TODO(), bson.M{"keyId": nodeId}, bson.M{"lastTime": latestTime}, update)
}

func NewNodeDao() dao.INodeDao {
	return &NodeDao{dao: mongodb.NewDaoWithTable(mongodb.NodeTable)}
}
