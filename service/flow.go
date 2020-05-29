package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/dao"
	"pervasive-chain/db"
	"pervasive-chain/form"
)

type FlowService struct {
	dao dao.IDao
}

func (f *FlowService) UpdateFlowInfo(flowForm form.ReportFlowForm) (interface{}, error) {
	param := bson.M{
		"type":   flowForm.Time,
		"number": flowForm.Number,
		"time":   flowForm.Time,
		"in":     flowForm.In,
		"out":    flowForm.Out,
	}
	update := options.Update()
	update.SetUpsert(true)
	return f.dao.UpdateWithOption(bson.M{"id": flowForm.Id}, bson.M{"$set": param}, update)
}

func NewFlowService() IFlowService {
	return &FlowService{dao: dao.NewDao(db.FlowTable)}
}
