package form

type ReportFlowForm struct {
	Type   string `form:"type" binding:"required" json:"type"`     //[b|r|s], 链类型
	Number string `form:"number" binding:"required" json:"number"` // 链编号
	Id     string `form:"id" binding:"required" json:"id"`         // 节点id
	Time   int64  `form:"time" binding:"required" json:"time"`     //当前产生时间
	In     int64  `form:"in" binding:"required" json:"in"`         //下行带宽
	Out    int64  `form:"out" binding:"required" json:"out"`       //上行带宽
}

func (h *ReportFlowForm) Valid() (bool, error) {
	panic(h)
}
