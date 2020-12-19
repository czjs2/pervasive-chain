package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/dao"
	"pervasive-chain/model"
	"pervasive-chain/mongodb"
	"time"
)

type NodeDao struct {
	dao mongodb.IDao
}

func (n *NodeDao) UpdateNodeCmd(chainType string, amount int) (interface{}, error) {
	return nil, n.dao.UseSession(context.TODO(), func(sessionContext context.Context) error {
		param := bson.M{
			"cmd":     bson.M{"key": "transfer", "params": bson.M{"amount": amount}},
			"cmdTime": time.Now(),
		}
		_, err := n.dao.UpdateMany(sessionContext, bson.M{"type": chainType}, param)
		if err != nil {
			return nil
		}
		return nil
	})
}

func (n *NodeDao) TotalNode(chainType string) (int, error) {
	total := model.Total{}
	query := []bson.M{
		bson.M{"$match": bson.M{"type": chainType}},
		bson.M{"$group": bson.M{"_id": "", "total": bson.M{"$sum": 1}}},
		bson.M{"$project": bson.M{"total": 1}},
	}
	return total.Total, n.dao.AggregateOne(context.TODO(), query, &total)
}

func (n *NodeDao) FindLatestOne(chainType string) (*model.Node, error) {
	obj := &model.Node{}
	query := []bson.M{
		bson.M{"$match": bson.M{"type": chainType}},
		bson.M{"$limit": 1},
	}
	return obj, n.dao.AggregateOne(context.TODO(), query, obj)
}

func (n *NodeDao) Insert(chainType, chainKey, nodeId string) (interface{}, error) {

	return n.dao.InsertOne(context.TODO(), bson.M{"type": chainType, "chainKey": chainKey, "nodeId": nodeId, "lastTime": time.Now()})

}

func (n *NodeDao) FindOne(nodeId string) (*model.Node, error) {
	obj := &model.Node{}
	_, err := n.dao.FindOne(context.TODO(), bson.M{"nodeId": nodeId}, obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (n *NodeDao) UpdateLatestTime(nodeId string) (interface{}, error) {
	update := options.Update()
	update.SetUpsert(true)
	return n.dao.UpdateWithOption(context.TODO(), bson.M{"nodeId": nodeId}, bson.M{"lastTime": time.Now()}, update)
}

func NewNodeDao() dao.INodeDao {
	return &NodeDao{dao: mongodb.NewDaoWithTable(mongodb.NodeTable)}
}
