package node

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/form"
	"pervasive-chain/utils"
)

func HeartBeatHandler(c *gin.Context) {
	var heartBeatFrom form.HeartBeatFrom
	err := c.BindJSON(&heartBeatFrom)
	if err != nil {
		utils.FailResponse(c)
		return
	}
	nodeService := NewNodeService()
	_, err = nodeService.UpdateNodeInfo(heartBeatFrom)
	if err != nil {
		utils.FailResponse(c)
		return
	}
	utils.SuccessResponse(c, nil)
}

type NodeService struct {
	nodeDao dao.INodeDao
}


func (n *NodeService) UpdateNodeInfo(heartFrom form.HeartBeatFrom) (interface{}, error) {
	//
	//keyId := fmt.Sprintf("%v%v", heartFrom.Type, heartFrom.Id)
	//node, err := n.nodeDao.FindOne(keyId)
	//if err != nil {
	//	return nil, fmt.Errorf("query node info error: %v  %v \n", keyId, err)
	//}
	//_, err = n.nodeDao.UpdateLatestTime(keyId, heartFrom.Time)
	//if err != nil {
	//	return nil, fmt.Errorf("update node info error %v %v \n", keyId, err)
	//}

	return nil,nil
}

func NewNodeService() INodeService {
	return &NodeService{nodeDao: daoimpl.NewNodeDao()}
}
