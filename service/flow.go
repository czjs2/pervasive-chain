package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/config"
	"pervasive-chain/dao"
	"pervasive-chain/db"
	"pervasive-chain/form"
	"pervasive-chain/model"
)

type FlowService struct {
	dao dao.IDao
}

func (f *FlowService) UpdateFlowInfo(flowForm form.ReportFlowForm) (interface{}, error) {
	nodeId := fmt.Sprintf("%s-%s", flowForm.Type, flowForm.Id)
	param := bson.M{
		"nodeId": nodeId,
		"time":   flowForm.Time,
		"in":     flowForm.In,
		"out":    flowForm.Out,
	}

	update := options.Update()
	update.SetUpsert(true)
	return f.dao.UpdateWithOption(bson.M{"nodeId": nodeId},  param, update)
}

func NewFlowService() IFlowService {
	return &FlowService{dao: dao.NewDao(db.FlowTable)}
}

type TotalFlowService struct {
	dao dao.IDao
}

func (t *TotalFlowService) AddTotalFlow(flowForm form.TotalFlowForm) (interface{}, error) {
	param := bson.M{
		"in":   flowForm.In,
		"out":  flowForm.Out,
		"time": flowForm.Time,
	}
	return t.dao.Add(param)
}

func (t *TotalFlowService) FlowList() (interface{}, int, error) {
	query := []bson.M{
		bson.M{"$sort": bson.M{"createTime": -1}},
		bson.M{"$limit": config.PageSize},
	}
	flow := model.TotalFlow{}
	return t.dao.List(query, &flow)
}

func NewTotalFlowService() ITotalFlowService {
	return &TotalFlowService{dao: dao.NewDao(db.TotalFlowTable)}
}
