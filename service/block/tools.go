package block

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"pervasive-chain/utils"
)

func getBlockParams(blockFrom ReportBlockForm) (bson.M, error) {
	params := bson.M{}
	detail := bson.M{}
	time, err := utils.ParseLocalTime(blockFrom.Time)
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

func getShardTransParam(blockFrom ReportBlockForm) ([]interface{}, []interface{}) {
	var group []RelayTransGroup
	group = append(blockFrom.Detail.UpStream)
	group = append(blockFrom.Detail.DownStream)
	var res []interface{}
	var trans []interface{}
	for i := 0; i < len(group); i++ {
		up := blockFrom.Detail.UpStream[i]
		for j := 0; j < len(blockFrom.Detail.Ss); j++ {
			ss := blockFrom.Detail.Ss[j]

			if fmt.Sprintf("%v-%v", ss.FromRelay, ss.ToRelay) == up.Key {
				res = append(res, bson.M{
					"fromShard": ss.FromShard,
					"toShard":   ss.ToShard,
					"fromRelay": ss.FromRelay,
					"toRelay":   ss.ToRelay,
					"hash":      up.Hash,
					"trans":     len(ss.Trans),
				})
			}

			// 交易数据
			for k := 0; k < len(ss.Trans); k++ {
				tran := ss.Trans[k]
				trans = append(trans, bson.M{
					"fromShard": ss.FromShard,
					"toShard":   ss.ToShard,
					"fromRelay": ss.FromRelay,
					"toRelay":   ss.ToRelay,
					"height":    blockFrom.Height,
					"hash":      up.Hash,
					"from":      tran.From,
					"to":        tran.To,
					//"amount":     0,
				})
			}

		}
	}
	return res,trans

}

