package form

type ReportBlockForm struct {
	Type     string      `form:"type" binding:"required" json:"type"`         //[b|r|s], 链类型
	Number   string      `form:"number" binding:"required" json:"number"`     // 链编号
	Id       string      `form:"id" binding:"required" json:"id"`             // 节点id
	Height   int64       `form:"height" binding:"required" json:"height"`     //当前区块高度
	Father   string      `form:"father" binding:"required" json:"father"`     //父区块hash
	Hash     string      `form:"hash" binding:"required" json:"hash"`         //区块hash
	Vrf      string      `form:"vrf" binding:"omitempty" json:"vrf"`           //VRF
	Time     int64       `form:"time" binding:"required" json:"time"`         //当前产生时间
	Interval int64       `form:"interval" binding:"required" json:"interval"` //出块间隔
	Trans    int64       `form:"trans" binding:"required" json:"trans"`       //交易数量
	Size     int64       `form:"size" binding:"required" json:"size"`         //区块大小
	Detail   interface{} `form:"detail" json:"detail"`                        //详情 (需详细定义)
}

func (h *ReportBlockForm) Valid() (bool, error) {
	panic(h)
}


