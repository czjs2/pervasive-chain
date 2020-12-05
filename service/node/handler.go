package node

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/form"
	"pervasive-chain/utils"
)

type NodeHandler struct {
	nodeDao dao.INodeDao
}

func (n *NodeHandler) UpdateNodeInfo(c *gin.Context) {
	fmt.Printf("heart beat ... %v  \n", c.Request.RequestURI)
	var heartFrom form.HeartBeatFrom
	utils.MustParams(c, &heartFrom)
	_, err := n.nodeDao.UpdateLatestTime(heartFrom.Id, heartFrom.Time)
	if err != nil {
		utils.FailResponse(c)
		return
	}
	utils.SuccessResponse(c, nil)
}

func NewNodeService() *NodeHandler {
	return &NodeHandler{nodeDao: daoimpl.NewNodeDao()}
}
