package block

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"pervasive-chain/config"
	"pervasive-chain/utils"
)

func getTransGroupParam(blockFrom ReportBlockForm) ([]interface{}, []interface{}) {
	switch blockFrom.Type {
	case config.BeaconType:
		return nil, nil
	case config.RelayType:
		return getRelayParam(blockFrom), nil
	case config.SharedType:
		return getShardParam(blockFrom)
	default:
		return nil, nil
	}
}

func getRelayParam(form ReportBlockForm) []interface{} {
	var transGroup []interface{}
	for i := 0; i < len(form.Detail.Ss); i++ {
		ss := form.Detail.Ss[i]
		transGroup = append(transGroup, bson.M{
			"hash":      ss.Hash,
			"height":    form.Height,
			"trans":     len(ss.Trans),
			"fromShard": ss.FromShard,
			"toShard":   ss.ToShard,
			"fromRelay": ss.FromRelay,
			"toRelay":   ss.ToRelay,
		})

	}
	return transGroup
}

func getShardParam(form ReportBlockForm) ([]interface{}, []interface{}) {
	var transGroup []interface{}
	var trans []interface{}
	for i := 0; i < len(form.Detail.Ss); i++ {
		ss := form.Detail.Ss[i]
		transGroup = append(transGroup, bson.M{
			"hash":      ss.Hash,
			"height":    form.Height,
			"trans":     len(ss.Trans),
			"fromShard": ss.FromShard,
			"toShard":   ss.ToShard,
			"fromRelay": ss.FromRelay,
			"toRelay":   ss.ToRelay,
		})
		for j := 0; j < len(ss.Trans); j++ {
			tran := ss.Trans[j]
			trans = append(trans, bson.M{
				"hash":      tran.Hash,
				"height":    form.Height,
				"from":      tran.From,
				"fromShard": ss.FromShard,
				"to":        tran.To,
				"toShard":   ss.ToShard,
				"amount":    tran.Amount,
			})
		}
	}
	return transGroup, trans
}

func getRealBlockParam(blockFrom ReportBlockForm) bson.M {
	param := bson.M{}
	param["nodeId"] = blockFrom.NodeId
	param["father"] = blockFrom.Father
	param["hash"] = blockFrom.Hash
	param["vrf"] = blockFrom.Vrf
	param["time"] = blockFrom.Time
	param["type"] = blockFrom.Type
	param["chainKey"] = blockFrom.ChainKey
	param["height"] = blockFrom.Height
	param["interval"] = blockFrom.Interval
	param["trans"] = blockFrom.Trans
	param["size"] = blockFrom.Size
	return param
}

func getLatestParams(blockFrom ReportBlockForm) (bson.M, error) {
	param := bson.M{}
	time, err := utils.ParseRFCTime(blockFrom.Time)
	if err != nil {
		return nil, err
	}
	param["time"] = time
	param["type"] = blockFrom.Type
	param["chainKey"] = blockFrom.ChainKey
	param["height"] = blockFrom.Height
	param["interval"] = blockFrom.Interval
	param["trans"] = blockFrom.Trans
	if blockFrom.Interval == 0.0 {
		param["tps"] = blockFrom.Trans
	}else {
		param["tps"] = float64(blockFrom.Trans) / blockFrom.Interval
	}

	param["size"] = blockFrom.Size
	return param, nil

}

func getBlockParams(blockFrom ReportBlockForm) (bson.M, error) {
	params := bson.M{}
	detail := bson.M{}
	time, err := utils.ParseRFCTime(blockFrom.Time)
	if err != nil {
		return nil, err
	}
	detail["upStream"] = blockFrom.Detail.UpStream
	detail["downStream"] = blockFrom.Detail.DownStream
	params["type"] = blockFrom.Type
	params["chainKey"] = blockFrom.ChainKey
	params["miner"] = blockFrom.NodeId
	params["height"] = blockFrom.Height
	params["father"] = blockFrom.Father
	params["hash"] = blockFrom.Hash
	params["vrf"] = blockFrom.Vrf
	params["time"] = time
	params["gas"] = blockFrom.Gas
	params["interval"] = blockFrom.Interval
	params["trans"] = blockFrom.Trans
	params["size"] = blockFrom.Size
	params["lockHash"] = blockFrom.LockHash
	params["upHash"] = blockFrom.UpHash
	params["downHash"] = blockFrom.DownHash
	params["detail"] = detail
	return params, nil
}

func getTransGroupParamV1(blockFrom ReportBlockForm) (interface{}, interface{}) {
	switch blockFrom.Type {
	case config.BeaconType:
		return nil, nil
	case config.RelayType:
		return getRelayParamV1(blockFrom), nil
	case config.SharedType:
		return getSharedParamV1(blockFrom)
	default:
		return nil, nil
	}
}

func getSharedParamV1(form ReportBlockForm) (interface{}, interface{}) {
	var transGroup []mongo.WriteModel
	var trans []mongo.WriteModel
	for i := 0; i < len(form.Detail.Ss); i++ {
		ss := form.Detail.Ss[i]
		UpdateManyModel := mongo.NewUpdateOneModel()
		UpdateManyModel.SetUpsert(true)
		UpdateManyModel.SetFilter(bson.M{"hash": ss.Hash})
		UpdateManyModel.SetUpdate(bson.M{"$set": bson.M{
			"trans":     len(ss.Trans),
			"height":    form.Height,
			"fromShard": ss.FromShard,
			"toShard":   ss.ToShard,
			"fromRelay": ss.FromRelay,
			"toRelay":   ss.ToRelay,
			"hash":      ss.Hash,
		}})
		transGroup = append(transGroup, UpdateManyModel)
		for j := 0; j < len(ss.Trans); j++ {
			tran := ss.Trans[j]
			transUpdateManyModel := mongo.NewUpdateOneModel()
			transUpdateManyModel.SetUpsert(true)
			transUpdateManyModel.SetFilter(bson.M{"hash": tran.Hash})
			transUpdateManyModel.SetUpdate(bson.M{"$set": bson.M{
				"hash":      tran.Hash,
				"height":    form.Height,
				"from":      tran.From,
				"fromShard": ss.FromShard,
				"to":        tran.To,
				"toShard":   ss.ToShard,
				"amount":    tran.Amount,
			}})
			trans = append(trans, transUpdateManyModel)
		}
	}
	return transGroup, trans
}

func getRelayParamV1(form ReportBlockForm) interface{} {
	var transGroup []mongo.WriteModel
	for i := 0; i < len(form.Detail.Ss); i++ {
		ss := form.Detail.Ss[i]
		UpdateManyModel := mongo.NewUpdateOneModel()
		UpdateManyModel.SetUpsert(true)
		UpdateManyModel.SetFilter(bson.M{"hash": ss.Hash})
		UpdateManyModel.SetUpdate(bson.M{
			"trans":     len(ss.Trans),
			"height":    form.Height,
			"fromShard": ss.FromShard,
			"toShard":   ss.ToShard,
			"fromRelay": ss.FromRelay,
			"toRelay":   ss.ToRelay,
			"hash":      ss.Hash,
		})
		transGroup = append(transGroup, UpdateManyModel)
	}
	return transGroup
}
