package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/config"
	"pervasive-chain/db"
	_ "pervasive-chain/db"
	"pervasive-chain/model"
	"pervasive-chain/utils"
)

type StatisticService struct {
}

func (s *StatisticService) CountFlow() (interface{}, error) {
	// todo 需要强一致性?
	collection := db.Collection(db.FlowTable)
	cursor, err := collection.Aggregate(context.TODO(), []bson.M{
		bson.M{"$group": bson.M{"_id": "", "out": bson.M{"$sum": "$out"}, "in": bson.M{"$sum": "$in"}}},
	})
	defer db.CloseCursor(cursor)
	if err != nil {
		return nil, nil
	}
	totalFlowCollection := db.Collection(db.TotalFlowTable)
	flow := model.TotalFlow{}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&flow)
		if err != nil {
			return nil, err
		}
	}
	_, err = totalFlowCollection.InsertOne(context.TODO(), bson.M{"out": flow.Out, "in": flow.In, "createTime": utils.GetNowTime()})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *StatisticService) CountChain() (interface{}, error) {
	// todo 需要强一致性？
	collection := db.Collection(db.Node)
	cursor, err := collection.Aggregate(context.TODO(), []bson.M{
		bson.M{"$group": bson.M{"_id": "$type", "total": bson.M{"$sum": 1}}},
	})
	defer db.CloseCursor(cursor)
	if err != nil {
		return nil, err
	}
	var res []model.ChainType
	for cursor.Next(context.TODO()) {
		chain := model.ChainType{}
		err := cursor.Decode(&chain)
		if err != nil {
			return nil, err
		}
		res = append(res, chain)
	}
	totalChainInfo := model.TotalChain{}
	sum := 0
	for i := 0; i < len(res); i++ {
		if res[i].Id == config.BChain {
			totalChainInfo.RelayNum = res[i].Total
		} else if res[i].Id == config.SChain {
			totalChainInfo.SharedNum = res[i].Total
		}
		sum = sum + res[i].Total
	}
	totalChainInfo.NodeNum = sum
	// 总能力
	blockCollection := db.Collection(db.BlockInfoTable)
	cursor1, err := blockCollection.Aggregate(context.TODO(), []bson.M{
		// todo 统计时间维度
		bson.M{"$match": bson.M{"time": bson.M{"$gte": "", "$lte": ""}}},
		bson.M{"$group": bson.M{"_id": "", "tps": bson.M{"$avg": "$trans"}}},
	})
	defer db.CloseCursor(cursor1)
	if err != nil {
		return nil, err
	}
	tps := model.ChainTps{}
	for cursor1.Next(context.TODO()) {
		err := cursor1.Decode(&tps)
		if err != nil {
			return nil, err
		}
	}
	totalChainInfo.TotalNum = tps.Tps
	totalChainCollection := db.Collection(db.TotalChainTable)
	update := options.FindOneAndUpdate()
	update.SetUpsert(true)
	_, err = totalChainCollection.InsertOne(context.TODO(), bson.M{"relayNum": totalChainInfo.RelayNum, "sharedNum": totalChainInfo.SharedNum,
		"nodeNum": totalChainInfo.NodeNum, "totalNum": totalChainInfo.TotalNum,
		"createTime": utils.GetNowTime()})
	if err != nil {
		return nil, err
	}
	return nil, nil

}

func NewStatisticService() IStatisticsService {
	return &StatisticService{}
}
