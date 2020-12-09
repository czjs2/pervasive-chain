package daoimpl

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	model2 "pervasive-chain/model"
)

func getQueryBlockParam(chainType, chainKey, hash string, height int) bson.M {
	query := bson.M{}
	if hash != "" {
		query["hash"] = hash
		return query
	}
	if chainType != "" || chainKey != "" || height != 0 {
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
