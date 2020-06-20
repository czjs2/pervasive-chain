package service

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/config"
	"pervasive-chain/dao"
	"pervasive-chain/db"
	"pervasive-chain/form"
)

type FlowService struct {
	dao dao.IDao
}

func (f *FlowService) UpdateFlowInfo(flowForm form.ReportFlowForm) (interface{}, error) {
	nodeId := fmt.Sprintf("%s-%s", flowForm.Type, flowForm.Id)
	if flowForm.Time<=0{
		return nil,errors.New("time is zero")
	}
	param := bson.M{
		"nodeId": nodeId,
		"time":   millisecondToTime(flowForm.Time),
		"in":     flowForm.In,
		"out":    flowForm.Out,
	}

	update := options.Update()
	update.SetUpsert(true)
	return f.dao.UpdateWithOption(bson.M{"nodeId": nodeId}, param, update)
}

func NewFlowService() IFlowService {
	return &FlowService{dao: dao.NewDao(db.FlowTable)}
}

type TotalFlowService struct {
	dao dao.IDao
}

func (t *TotalFlowService) AddTotalFlow(flowForm form.TotalFlowForm) (interface{}, error) {
	if flowForm.Time<=0{
		return nil,errors.New("time is zero")
	}
	param := bson.M{
		"in":   flowForm.In,
		"out":  flowForm.Out,
		"time": millisecondToTime(flowForm.Time),
	}
	return t.dao.Add(param)
}

func (t *TotalFlowService) FlowList() (interface{}, int, error) {
	query := []bson.M{
		bson.M{"$sort": bson.M{"time": -1}},
		bson.M{"$limit": config.PageSize},
	}
	return t.dao.List(query)
}

func NewTotalFlowService() ITotalFlowService {
	return &TotalFlowService{dao: dao.NewDao(db.TotalFlowTable)}
}
