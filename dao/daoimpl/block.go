package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/dao"
	"pervasive-chain/mongodb"
)

type BlockDao struct {
	dao        mongodb.IDao
	trans      mongodb.IDao
	transGroup mongodb.IDao
}

func (b *BlockDao) Insert(blockParam bson.M, transGroup, trans []interface{}) (interface{}, error) {
	update := options.Update()
	update.SetUpsert(true)
	query := bson.M{"hash": blockParam["hash"], "height": blockParam["height"], "chainKey": blockParam["chainKey"]}
	_, err := b.dao.UpdateWithOption(context.TODO(), query, bson.M(blockParam), update)
	if err!=nil{
		return nil,err
	}
	// todo 更优的方式
	_, err = b.transGroup.InsertMany(context.TODO(), transGroup)
	if err!=nil{
		return nil,err
	}
	_, err = b.trans.InsertMany(context.TODO(), trans)
	if err!=nil{
		return nil,err
	}
	return nil,nil

}

func (b *BlockDao) Query() (interface{}, error) {
	panic("implement me")
}

func NewBlockDao() dao.IBlockDao {
	return &BlockDao{
		dao: mongodb.NewDaoWithTable(mongodb.BlocksTable),
		trans:mongodb.NewDaoWithTable(mongodb.TransTable),
		transGroup:mongodb.NewDaoWithTable(mongodb.GroupTable),
	}
}
