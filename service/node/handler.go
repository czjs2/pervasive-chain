package node

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/config"
	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/statecode"
	"pervasive-chain/utils"
	"pervasive-chain/ws"
	"time"
)

type NodeHandler struct {
	nodeDao dao.INodeDao
}

func (n *NodeHandler) GenCmd(c *ws.WsContext) {
	var genCmdFrom GenCmdFrom
	_ = c.BindJSON(&genCmdFrom)
	node, err := n.nodeDao.FindLatestOne()
	if err != nil {
		utils.WsFailResponse(c)
		return
	}
	if node.NodeId == "" {
		utils.WsFailResponse(c)
		return
	}
	if time.Now().Sub(node.CmdTime).Seconds() > config.GenCmdIntervalTime {
		utils.WsFailResponse(c)
		return
	}
	utils.WsSuccessResponse(c, nil)

}

func (n *NodeHandler) UpdateNodeInfo(c *gin.Context) {
	var heartFrom HeartBeatFrom
	utils.MustParams(c, &heartFrom)
	node, err := n.nodeDao.FindOne(heartFrom.NodeId)
	if err != nil {
		if err.Error() != statecode.NoResultErr {
			utils.FailResponse(c)
			return
		}
	}
	if node == nil {
		_, err := n.nodeDao.Insert(heartFrom.Type, heartFrom.ChainKey, heartFrom.NodeId, heartFrom.Time)
		if err != nil {
			utils.FailResponse(c)
			return
		}
	} else {
		_, err := n.nodeDao.UpdateLatestTime(heartFrom.NodeId, heartFrom.Time)
		if err != nil {
			utils.FailResponse(c)
			return
		}
		if !node.CmdTime.IsZero() && time.Now().UTC().Sub(node.CmdTime) > config.GenCmdIntervalTime {
			utils.SuccessResponse(c, node.Cmd)
			return
		}
	}
	utils.SuccessResponse(c, nil)

}

func NewNodeService() *NodeHandler {
	return &NodeHandler{nodeDao: daoimpl.NewNodeDao()}
}
