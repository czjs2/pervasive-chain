package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/dao"
	"pervasive-chain/model"
	"pervasive-chain/mongodb"
)

type BlockDao struct {
	dao        mongodb.IDao
	trans      mongodb.IDao
	transGroup mongodb.IDao
	realBlock  mongodb.IDao
}

func (b *BlockDao) Block(chainType, chainKey, hash string, height string) (interface{}, error) {
	param := &model.Param{}
	err := b.dao.AggregateOne(context.TODO(), []bson.M{}, param)
	if err != nil {
		return nil, err
	}
	// todo 改表 多次查询？
	return nil, err
}

func (b *BlockDao) Insert(blockParam, latestParam bson.M, transGroup, trans []interface{}) (interface{}, error) {
	update := options.Update()
	update.SetUpsert(true)
	err := b.dao.UseSession(context.TODO(), func(sessionContext context.Context) error {
		query := bson.M{"hash": blockParam["hash"], "height": blockParam["height"], "chainKey": blockParam["chainKey"]}
		_, err := b.dao.UpdateWithOption(sessionContext, query, bson.M(blockParam), update)
		if err != nil {
			return err
		}
		_, err = b.realBlock.UpdateWithOption(sessionContext, bson.M{"chainKey": latestParam["chainKey"]}, latestParam, update)
		if err != nil {
			return err
		}
		// todo 更优的方式
		if len(transGroup) > 0 {
			_, err = b.transGroup.InsertMany(sessionContext, transGroup)
			if err != nil {
				return err
			}
		}
		if len(trans) > 0 {
			_, err = b.trans.InsertMany(sessionContext, trans)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil, err

}

func (b *BlockDao) Query() (interface{}, error) {
	panic("implement me")
}

func NewBlockDao() dao.IBlockDao {
	return &BlockDao{
		dao:        mongodb.NewDaoWithTable(mongodb.BlocksTable),
		trans:      mongodb.NewDaoWithTable(mongodb.TransactionsTable),
		transGroup: mongodb.NewDaoWithTable(mongodb.TransGroupsTable),
		realBlock:  mongodb.NewDaoWithTable(mongodb.RealChainInfosTable),
	}
}
