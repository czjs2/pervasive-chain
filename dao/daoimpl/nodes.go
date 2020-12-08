package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/config"
	"pervasive-chain/dao"
	"pervasive-chain/model"
	"pervasive-chain/mongodb"
	"pervasive-chain/utils"
)

type NodeDao struct {
	dao mongodb.IDao
}

func (n *NodeDao) FindLatestOne() (*model.Node, error) {
	obj := &model.Node{}
	query := []bson.M{
		bson.M{"$match": bson.M{"type": config.SharedType}},
		bson.M{"$limit": 1},
	}
	return obj, n.dao.AggregateOne(context.TODO(), query, obj)
}

func (n *NodeDao) Insert(chainType, chainKey, nodeId, latestTime string) (interface{}, error) {
	time, err := utils.ParseRFCTime(latestTime)
	if err != nil {
		return nil, err
	}
	return n.dao.InsertOne(context.TODO(), bson.M{"type": chainType, "chainKey": chainKey, "nodeId": nodeId, "lastTime": time})

}

func (n *NodeDao) FindOne(nodeId string) (*model.Node, error) {
	obj := &model.Node{}
	_, err := n.dao.FindOne(context.TODO(), bson.M{"nodeId": nodeId}, obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (n *NodeDao) UpdateLatestTime(nodeId, latestTime string) (interface{}, error) {
	update := options.Update()
	update.SetUpsert(true)
	time, err := utils.ParseRFCTime(latestTime)
	if err != nil {
		return nil, err
	}
	return n.dao.UpdateWithOption(context.TODO(), bson.M{"nodeId": nodeId}, bson.M{"lastTime": time}, update)
}

func NewNodeDao() dao.INodeDao {
	return &NodeDao{dao: mongodb.NewDaoWithTable(mongodb.NodeTable)}
}
