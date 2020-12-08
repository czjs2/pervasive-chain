package block

import (
	"go.mongodb.org/mongo-driver/bson"
	"pervasive-chain/utils"
)

func getShardTransParam(blockFrom ReportBlockForm) ([]interface{}, []interface{}) {
	var group []RelayTransGroup
	group = append(blockFrom.Detail.UpStream)
	group = append(blockFrom.Detail.DownStream)
	var res []interface{}
	var trans []interface{}
	for i := 0; i < len(group); i++ {
		for j := 0; j < len(blockFrom.Detail.Ss); j++ {
			ss := blockFrom.Detail.Ss[j]

			// 跨分片交易分组
			res = append(res, bson.M{
				"fromShard": ss.FromShard,
				"toShard":   ss.ToShard,
				"fromRelay": ss.FromRelay,
				"toRelay":   ss.ToRelay,
				"hash":      ss.Hash,
				"trans":     len(ss.Trans),
				"height":    blockFrom.Height,
			})

			// 交易数据
			for k := 0; k < len(ss.Trans); k++ {
				tran := ss.Trans[k]
				trans = append(trans, bson.M{
					"fromShard": ss.FromShard,
					"toShard":   ss.ToShard,
					"fromRelay": ss.FromRelay,
					"toRelay":   ss.ToRelay,
					"height":    blockFrom.Height,
					"hash":      ss.Hash, // todo 交易hash
					"from":      tran.From,
					"to":        tran.To,
					"amount":    tran.Amount,
				})
			}

		}
	}
	return res, trans

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
	param["tps"] = blockFrom.Trans
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
	params["nodeId"] = blockFrom.NodeId
	params["height"] = blockFrom.Height
	params["father"] = blockFrom.Father
	params["hash"] = blockFrom.Hash
	params["vrf"] = blockFrom.Vrf
	params["time"] = time
	params["interval"] = blockFrom.Interval
	params["trans"] = blockFrom.Trans
	params["size"] = blockFrom.Size
	params["lockHash"] = blockFrom.LockHash
	params["upHash"] = blockFrom.UpHash
	params["downHash"] = blockFrom.DownHash
	params["detail"] = detail
	return params, nil
}
