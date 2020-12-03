package form

type HeartBeatFrom struct {
	Type   string `form:"type" binding:"required" json:"type"`     //[b|r|s], 链类型
	Number string `form:"number" binding:"required" json:"number"` // 链编号
	Id     string `form:"id" binding:"required" json:"id"`         // 节点id
	Time   int64  `form:"time" binding:"required" json:"time"`     // 时间

}

func (h *HeartBeatFrom) Valid() (bool, error) {
	panic(h)
}
