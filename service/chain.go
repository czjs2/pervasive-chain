package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/config"
	"pervasive-chain/dao"
	"pervasive-chain/db"
	"pervasive-chain/form"
	"pervasive-chain/model"
)

type ChainService struct {
	dao dao.IDao
}

func (s *ChainService) Chain(chainId string) (interface{}, error) {
	var query []bson.M
	if chainId == "" {
		// todo  返回信标链 ?
		query = append(query, bson.M{"$match": bson.M{}})
	} else {
		query = append(query, bson.M{"$match": bson.M{"id": chainId}})
	}

	chain := model.TotalChain{}

	return s.dao.Aggregate(query, &chain)
}

func (s *ChainService) ChainList() (interface{}, int, error) {
	qury := []bson.M{
		bson.M{"$sort": bson.M{"updateTime": -1}},
		bson.M{"$limit": config.PageSize},
	}
	totalShared := model.TotalChain{}
	return s.dao.List(qury, &totalShared)
}

func (s *ChainService) UpdateSharedInfo(shardForm form.ShardInfoForm) (interface{}, error) {
	param := bson.M{
		"relayNum":  shardForm.RelayNum,
		"sharedNum": shardForm.SharedNum,
		"nodeNum":   shardForm.NodeNum,
		"totalNum":  shardForm.TotalNum,
	}
	return s.dao.Add(param)
}

func (s *ChainService) LatestShardInfo() (interface{}, error) {
	query := []bson.M{
		bson.M{"$sort": bson.M{"updateTime": -1}},
		bson.M{"$limit": 1},
	}
	chainTotalInfo := model.TotalChain{}
	return s.dao.Aggregate(query, &chainTotalInfo)
}

func NewChainService() IChainService {
	return &ChainService{dao: dao.NewDao(db.TotalChainTable)}
}

type TotalChainService struct {
	dao dao.IDao
}


func (t *TotalChainService) UpdateTotalChainInfo(relayNum, sharedNum, nodeNum, tps int) (interface{}, error) {
	query := bson.M{
		"tps": bson.M{"$exists": true},
	}
	params := bson.M{
		"relayNum":  relayNum,
		"sharedNum": sharedNum,
		"nodeNum":   nodeNum,
		"tps":       tps,
	}
	update := options.Update()
	update.SetUpsert(true)
	return t.dao.UpdateWithOption(query,  params, update)

}

func (t *TotalChainService) TotalFlowList() (interface{}, int, error) {
	query := []bson.M{
		bson.M{"$sort": bson.M{"updateTime": -1}},
		bson.M{"$limit": 1},
	}
	chainTotalInfo := model.TotalChain{}
	return t.dao.List(query, &chainTotalInfo)
}

func NewTotalChainService() ITotalChainService {
	return &TotalChainService{dao: dao.NewDao(db.TotalChainTable)}
}
