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
	"time"
)

type NodeDao struct {
	dao mongodb.IDao
}

func (n *NodeDao) UpdateNodeCmd(amount int) (interface{}, error) {
	return nil, n.dao.UseSession(context.TODO(), func(sessionContext context.Context) error {
		param := bson.M{
			"cmd":     bson.M{"key": "transfer", "params": bson.M{"amount": amount}},
			"cmdTime": time.Now(),
		}
		_, err := n.dao.UpdateMany(sessionContext, bson.M{"type": config.SharedType}, param)
		if err != nil {
			return nil
		}
		// todo 效验
		return nil
	})
}

func (n *NodeDao) TotalShardNode() (int, error) {
	total:=model.Total{}
	query := []bson.M{
		bson.M{"$match": bson.M{"type": config.SharedType}},
		bson.M{"$group": bson.M{"_id": "", "total": bson.M{"$sum": 1}}},
		bson.M{"$project": bson.M{"total": 1}},
	}
	return total.Total, n.dao.AggregateOne(context.TODO(), query, &total)
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
	lTime, err := utils.ParseRFCTime(latestTime)
	if err != nil {
		return nil, err
	}
	return n.dao.InsertOne(context.TODO(), bson.M{"type": chainType, "chainKey": chainKey, "nodeId": nodeId, "lastTime": lTime})

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
	lTime, err := utils.ParseRFCTime(latestTime)
	if err != nil {
		return nil, err
	}
	return n.dao.UpdateWithOption(context.TODO(), bson.M{"nodeId": nodeId}, bson.M{"lastTime": lTime}, update)
}

func NewNodeDao() dao.INodeDao {
	return &NodeDao{dao: mongodb.NewDaoWithTable(mongodb.NodeTable)}
}
