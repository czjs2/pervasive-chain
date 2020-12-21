package daoimpl

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"pervasive-chain/config"
	model2 "pervasive-chain/model"
	"strings"
)

func getQueryTransGroup(fromShard, toShard string, height uint64) [] bson.M {
	match := bson.M{}
	if len(fromShard) == 3 {
		match["fromRelay"] = fromShard
	} else {
		match["fromShard"] = fromShard
	}
	if len(toShard) == 3 {
		match["toRelay"] = toShard
	} else {
		match["toShard"] = toShard
	}
	match["height"] = height
	query := []bson.M{
		bson.M{"$match": match},
		//bson.M{"$project": bson.M{"_id": 0, "hash": 1, "from": 1, "to": 1, "amount": 1}},
	}
	return query
}

func getQueryBlockParam(chainType, chainKey, hash string, height uint64) bson.M {
	query := bson.M{}
	if hash != "" {
		query["hash"] = hash
		return query
	}
	if strings.HasPrefix(chainType,config.BeaconType) {
		query["height"] = height
		query["type"] = chainType

	}else  {
		query["type"] = chainType
		query["chainKey"] = chainKey
		query["height"] = height
	}
	return query
}

func getTransGroup(res []interface{}) []mongo.WriteModel {
	var models []mongo.WriteModel
	for i := 0; i < len(res); i++ {
		param := res[i].(model2.Param)
		model := mongo.NewUpdateManyModel()
		model.SetUpsert(true)
		model.SetFilter(bson.M{"hash": param["hash"], "height": param["height"]})
		model.SetUpdate(bson.M{
			"trans":     param["trans"],
			"height":    param["height"],
			"fromShard": param["fromShard"],
			"toShard":   param["toShard"],
			"fromRelay": param["fromRelay"],
			"toRelay":   param["toRelay"],
			"hash":      param["hash"],
		})
		models = append(models, model)
	}
	return models
}
