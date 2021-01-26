package model

import "time"

type LatestBlock struct {
	Type     string    `json:"type" bson:"type"`
	ChainKey string    `json:"chainKey" bson:"chain_key"`
	Height   int       `json:"height" bson:"height"`
	Time     time.Time `json:"time" bson:"time"`
	Interval int       `json:"interval" bson:"interval"`
	Trans    int       `json:"trans" bson:"trans"`
	Tps      int       `json:"tps" bson:"tps"`
	Size     int       `json:"size" bson:"size"`
}

type BlockTime struct {
	Time time.Time `json:"time" bson:"time"`
}
