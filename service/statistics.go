package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/config"
	"pervasive-chain/db"
	_ "pervasive-chain/db"
	"pervasive-chain/form"
	"pervasive-chain/model"
	"pervasive-chain/utils"
	"time"
)

type StatisticService struct {
}

/*

   "height" : NumberLong(1),
    "detail" : null,
    "father" : "fatherHash",
    "hash" : "hash",
    "id" : "100003",
    "interval" : NumberLong(68),
    "number" : "100003",
    "size" : NumberLong(6068),
    "time" : NumberLong(1591432459),
    "trans" : NumberLong(6068),
    "type" : "s",
    "updateTime" : ISODate("2020-06-06T08:34:19.206+0000"),
    "vrf" : "vrf"

*/
func (s *StatisticService) AllChain() (interface{}, error) {
	// todo 暂时没找到更好的方式？
	collection := db.Collection(db.BlockInfoTable)
	var rootChain []interface{}
	cursor1, err := collection.Aggregate(context.TODO(), []bson.M{
		bson.M{"$group": bson.M{"_id": bson.M{"id": "$id", "type": "$type"}}},
	})
	defer db.CloseCursor(cursor1)
	if err != nil {
		return nil, err
	}
	var res []*model.Block1
	for cursor1.Next(context.TODO()) {
		ct := model.Block1{}
		err := cursor1.Decode(&ct)
		if err != nil {
			return nil, err
		}
		res = append(res, &ct)
	}
	for i := 0; i < len(res); i++ {
		cursor, err := collection.Aggregate(context.TODO(), []bson.M{
			bson.M{"$match": bson.M{"id": res[i].Id.Id}},
			bson.M{"$sort": bson.M{"height": -1}},
			bson.M{"$limit": 2},
		})
		defer db.CloseCursor(cursor)
		if err != nil {
			return nil, err
		}
		var blockList []*model.Block
		for cursor.Next(context.TODO()) {
			block := model.Block{}
			err := cursor.Decode(&block)
			if err != nil {
				return nil, err
			}
			blockList = append(blockList, &block)
		}
		param := make(map[string]interface{})
		param["blockList"] = blockList
		param["chainId"] = res[i].Id
		rootChain = append(rootChain, param)
	}
	return rootChain, nil
}

func (s *StatisticService) CountTps() (interface{}, error) {
	collection := db.Collection(db.TotalChainTable)
	cursor, err := collection.Aggregate(context.TODO(), []bson.M{
		bson.M{"$group": bson.M{"_id": "$timestamp", "tps": bson.M{"$sum": "$tps"}}},
	})
	defer db.CloseCursor(cursor)
	if err != nil {
		return nil, err
	}
	var res []*model.TotalTps
	for cursor.Next(context.TODO()) {
		tps := model.TotalTps{}
		err := cursor.Decode(&tps)
		if err != nil {
			return nil, err
		}
		res = append(res, &tps)
	}
	return res, nil
}

func (s *StatisticService) CountNode() (interface{}, error) {
	collection := db.Collection(db.Node)
	curTime := time.Now().Add(-config.NodeOffLineTime * time.Second)
	cursor, err := collection.Aggregate(context.TODO(), []bson.M{
		bson.M{"$match": bson.M{"lastTime": bson.M{"$gte": curTime.Unix()}}}, // 有效节点没有掉线
		bson.M{"$group": bson.M{"_id": "$type", "total": bson.M{"$sum": 1}}},
	})
	defer db.CloseCursor(cursor)
	if err != nil {
		return nil, err
	}
	var res []model.ChainType
	for cursor.Next(context.TODO()) {
		ct := model.ChainType{}
		err := cursor.Decode(&ct)
		if err != nil {
			return nil, err
		}
		res = append(res, ct)
	}
	return res, nil
}

func (s *StatisticService) CountFlow(flowForm form.ReportFlowForm) (interface{}, error) {
	totalFlowCollection := db.Collection(db.TotalFlowTable)
	update := options.FindOneAndUpdate()
	update.SetUpsert(true)
	flow := model.TotalFlow{}
	err := totalFlowCollection.FindOneAndUpdate(context.TODO(), bson.M{"time": flowForm.Time},
		bson.M{"$inc": bson.M{"out": flowForm.Out, "in": flowForm.In, "time": flowForm.Time}}, update).Decode(&flow)
	if err != nil {
		return nil, err
	}
	return flow, nil
}

func (s *StatisticService) CountChain(chainId, chainType string) (interface{}, error) {
	collection := db.Collection(db.BlockInfoTable)
	curStartTime, curEndTime, err := getStartAndEndTime()
	if err != nil {
		return nil, err
	}
	cursor, err := collection.Aggregate(context.TODO(), []bson.M{
		bson.M{"$match": bson.M{"number": chainId, "time": bson.M{"$gte": curStartTime.Unix(), "$lt": curEndTime.Unix()}}},
		bson.M{"$group": bson.M{"_id": "", "tps": bson.M{"$sum": "$tps"}}},
	})
	defer db.CloseCursor(cursor)
	if err != nil {
		return nil, err
	}
	tps := model.Tps{}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&tps)
		if err != nil {
			return nil, err
		}
	}
	totalChainCollection := db.Collection(db.TotalChainTable)
	update := options.FindOneAndUpdate()
	update.SetUpsert(true)
	result := totalChainCollection.FindOneAndUpdate(context.TODO(), bson.M{"timestamp": curStartTime.Unix(), "chainId": chainId},
		bson.M{"$set": bson.M{"chainId": chainId, "tps": tps.Tps, "type": chainType}}, update)
	if result.Err() != nil {
		return nil, err
	}
	return nil, nil

}

func getStartAndEndTime() (time.Time, time.Time, error) {
	curTime, err := utils.GetCurZeroTime()
	if err != nil {
		return curTime, curTime, err
	}
	tomTime := curTime.Add(24 * time.Hour)
	return curTime, tomTime, nil
}

func NewStatisticService() IStatisticsService {
	return &StatisticService{}
}
