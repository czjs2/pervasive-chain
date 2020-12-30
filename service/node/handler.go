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
	node, err := n.nodeDao.FindLatestOne(genCmdFrom.Type)
	if err != nil {
		utils.WsFailResponse(c)
		return
	}
	if node.NodeId == "" {
		utils.WsFailResponseWithMsg(c, "node num is zero,wait node report")
		return
	}
	if !node.CmdTime.IsZero() && time.Now().Sub(node.CmdTime).Seconds() < config.GenCmdIntervalTime {
		utils.WsFailResponseWithMsg(c, "previous gen is running,please send later")
		return
	}
	totalNode, err := n.nodeDao.TotalNode(genCmdFrom.Type)
	if err != nil {
		utils.WsFailResponse(c)
		return
	}
	if totalNode == 0 {
		utils.WsSuccessResponse(c, nil)
		return
	}
	_, err = n.nodeDao.UpdateNodeCmd(genCmdFrom.Type, genCmdFrom.Cmd.Params.Amount/totalNode)
	if err != nil {
		utils.WsFailResponse(c)
		return
	}
	utils.WsSuccessResponse(c, nil)

}

func (n *NodeHandler) UpdateNodeInfo(c *gin.Context) {
	var heartFrom HeartBeatFrom
	err := c.BindJSON(&heartFrom)
	if err != nil {
		utils.FailResponse(c, err.Error())
		return
	}
	if ok, err := heartFrom.Valid(); err != nil || !ok {
		utils.FailResponse(c, err)
		return
	}
	node, err := n.nodeDao.FindOne(heartFrom.NodeId)
	if err != nil {
		if err.Error() != statecode.NoResultErr {
			utils.FailResponse(c, err.Error())
			return
		}
	}
	if node == nil {
		_, err := n.nodeDao.Insert(heartFrom.Type, heartFrom.ChainKey, heartFrom.NodeId)
		if err != nil {
			utils.FailResponse(c, err.Error())
			return
		}
	} else {
		_, err := n.nodeDao.UpdateLatestTime(heartFrom.NodeId)
		if err != nil {
			utils.FailResponse(c, err.Error())
			return
		}
		// 两次下发命令的时间有间隔限制
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
