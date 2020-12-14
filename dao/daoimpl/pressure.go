package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (b *BlockDao) InsertV2(blockParam, latestParam bson.M, transGroup, trans [] interface{}) (interface{}, error) {
	update := options.Update()
	update.SetUpsert(true)
	err := b.dao.UseSession(context.TODO(), func(sessionContext context.Context) error {
		_, err := b.dao.InsertOne(sessionContext, bson.M(blockParam))
		if err != nil {
			return err
		}
		_, err = b.realBlock.InsertOne(sessionContext, latestParam)
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

func (b *BlockDao) InsertV3(blockParam, latestParam bson.M, transGroup, trans [] interface{}) (interface{}, error) {
	update := options.Update()
	update.SetUpsert(true)

	_, err := b.dao.InsertOne(context.TODO(), bson.M(blockParam))
	if err != nil {
		return nil, err
	}
	_, err = b.realBlock.InsertOne(context.TODO(), latestParam)
	if err != nil {
		return nil, err
	}
	// todo 更优的方式
	if len(transGroup) > 0 {
		_, err = b.transGroup.InsertMany(context.TODO(), transGroup)
		if err != nil {
			return nil, err
		}
	}
	if len(trans) > 0 {
		_, err = b.trans.InsertMany(context.TODO(), trans)
		if err != nil {
			return nil, err
		}
	}

	return nil, err
}
