package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/url"
	"os"
	"os/signal"
	"pervasive-chain/config"
	"pervasive-chain/db"
	"pervasive-chain/form"
	"pervasive-chain/utils"
	"time"
)

var addr = flag.String("addr", "192.168.0.164:9999", "http service address")
var host = "http://192.168.0.164:9999"

func main() {
	flag.Parse()
	interrupt := make(chan os.Signal, 1)
	interrupt1 := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt1, os.Interrupt)
//	go apiClient(interrupt1)
	time.Sleep(7 * time.Second)
	go socketClient(interrupt)
	select {}
}





func apiClient(interrupt chan os.Signal) {

	cleanDbData()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			reportData()
		case <-interrupt:
			return
		}
	}

}

var height1 int64 = 1
var height2 int64 = 1
var height3 int64 = 1

func reportData() {
	err := reportHeartBeat()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = reportBlock()
	if err != nil {
		fmt.Println(err.Error())
	}
	reportFlow()
}

func reportFlow() {
	reportFlowForm := form.ReportFlowForm{
		Type:   "s",
		Number: fmt.Sprintf("10000%d", utils.Rand(5)),
		Id:     fmt.Sprintf("10000%d", utils.Rand(5)),
		Time:   utils.GetTime().UnixNano(),
		In:     utils.Rand(1000),
		Out:    utils.Rand(1000),
	}
	_, err := ReportFlow(host, "/v1.0/flow", "", reportFlowForm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reportFlowForm2 := form.ReportFlowForm{
		Type:   "r",
		Number: fmt.Sprintf("10000%d", utils.Rand(5)),
		Id:     fmt.Sprintf("10000%d", utils.Rand(5)),
		Time:   utils.GetTime().UnixNano(),
		In:     utils.Rand(1000),
		Out:    utils.Rand(1000),
	}
	_, err = ReportFlow(host, "/v1.0/flow", "", reportFlowForm2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reportFlowForm3 := form.ReportFlowForm{
		Type:   "b",
		Number: fmt.Sprintf("10000%d", utils.Rand(5)),
		Id:     fmt.Sprintf("10000%d", utils.Rand(5)),
		Time:   utils.GetTime().UnixNano(),
		In:     utils.Rand(1000),
		Out:    utils.Rand(1000),
	}
	_, err = ReportFlow(host, "/v1.0/flow", "", reportFlowForm3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func reportBlock() error {
	reportBlockForm := form.ReportBlockForm{
		Type:     "s",
		Number:   fmt.Sprintf("10000%d", utils.Rand(5)),
		Id:       fmt.Sprintf("10000%d", utils.Rand(5)),
		Height:   height1,
		Father:   "fatherHash",
		Hash:     "hash",
		Vrf:      "vrf",
		Time:     utils.GetTime().UnixNano(),
		Interval: utils.Rand(100),
		Trans:    utils.Rand(10000),
		Size:     utils.Rand(10000),
		Detail:   nil,
	}
	height1 = height1 + 1
	_, err := ReportBlock(host, "/v1.0/block", "", reportBlockForm)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	reportBlockForm2 := form.ReportBlockForm{
		Type:     "r",
		Number:   fmt.Sprintf("10000%d", utils.Rand(5)),
		Id:       fmt.Sprintf("10000%d", utils.Rand(5)),
		Height:   height2,
		Father:   "fatherHash",
		Hash:     "hash",
		Vrf:      "vrf",
		Time:     utils.GetTime().UnixNano(),
		Interval: utils.Rand(100),
		Trans:    utils.Rand(10000),
		Size:     utils.Rand(10000),
		Detail:   nil,
	}
	height2 = height2 + 1
	_, err = ReportBlock(host, "/v1.0/block", "", reportBlockForm2)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	reportBlockForm3 := form.ReportBlockForm{
		Type:     "b",
		Number:   fmt.Sprintf("10000%d", utils.Rand(5)),
		Id:       fmt.Sprintf("10000%d", utils.Rand(5)),
		Height:   height3,
		Father:   "fatherHash",
		Hash:     "hash",
		Vrf:      "vrf",
		Time:     utils.GetTime().UnixNano(),
		Interval: utils.Rand(100),
		Trans:    utils.Rand(10000),
		Size:     utils.Rand(10000),
		Detail:   nil,
	}
	height3 = height3 + 1
	_, err = ReportBlock(host, "/v1.0/block", "", reportBlockForm3)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return err
}

func reportHeartBeat() error {
	bHeartform := form.HeartBeatFrom{
		Type:   "b",
		Number: fmt.Sprintf("%d", utils.Rand(1)),
		Id:     fmt.Sprintf("%d", utils.Rand(1)),
		Time:   utils.GetTime().UnixNano(),
	}
	rHeartform := form.HeartBeatFrom{
		Type:   "r",
		Number: fmt.Sprintf("%d", utils.Rand(1)),
		Id:     fmt.Sprintf("%d", utils.Rand(1)),
		Time:   utils.GetTime().UnixNano(),
	}
	sHeartform := form.HeartBeatFrom{
		Type:   "s",
		Number: fmt.Sprintf("%d", utils.Rand(1)),
		Id:     fmt.Sprintf("%d", utils.Rand(1)),
		Time:   utils.GetTime().UnixNano(),
	}
	_, err := HeartBeat(host, "/v1.0/heartbeat", "", bHeartform)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	_, err = HeartBeat(host, "/v1.0/heartbeat", "", rHeartform)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	_, err = HeartBeat(host, "/v1.0/heartbeat", "", sHeartform)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return err
}

func socketClient(interrupt chan os.Signal) {
	log.SetFlags(0)
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/v1.0/conn"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	done := make(chan struct{})

	defer c.Close()
	go func() {
		for {
			select {
			case <-interrupt:
				close(done)
				fmt.Println(" websocket interrupt")
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("write close:", err)
					return
				}
				return
			default:
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					return
				}
				log.Printf("recv: %s", message)
			}
		}
	}()

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			chainInfoCmd := fmt.Sprintf(`{"uri":"chainInfo","body":{},"msgId":"msgId%d"}`, time.Now().Unix())
			err = c.WriteMessage(websocket.TextMessage, []byte(chainInfoCmd))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			blockInfoCmd := fmt.Sprintf(`{"uri":"blockInfo","body":{"type":"b","number":"0"},"msgId":"msgId%d"}`, time.Now().Unix())
			err = c.WriteMessage(websocket.TextMessage, []byte(blockInfoCmd))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			cmdInfo := fmt.Sprintf(`{"uri":"cmd","body":{"type":"b","cmd":{"key":"transfer","params":[100]}},"msgId":"msgId%d"}`, time.Now().Unix())
			err := c.WriteMessage(websocket.TextMessage, []byte(cmdInfo))
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}


// just test
func cleanDbData() {

	prjConfig, err := config.ReadWebCfg("./web-config.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = db.InitMongo(prjConfig)

	collection := db.Collection(db.Node)
	collection1 := db.Collection(db.BlockInfoTable)
	collection2 := db.Collection(db.TotalChainTable)
	collection3 := db.Collection(db.HistoryBlockInfoTable)
	collection4 := db.Collection(db.NodeBandTable)
	collection5 := db.Collection(db.TotalBandWithTable)
	collection6 := db.Collection(db.FlowTable)
	collection7 := db.Collection(db.TotalFlowTable)
	_, _ = collection.DeleteMany(context.TODO(), bson.M{})
	_, _ = collection1.DeleteMany(context.TODO(), bson.M{})
	_, _ = collection2.DeleteMany(context.TODO(), bson.M{})
	_, _ = collection3.DeleteMany(context.TODO(), bson.M{})
	_, _ = collection4.DeleteMany(context.TODO(), bson.M{})
	_, _ = collection5.DeleteMany(context.TODO(), bson.M{})
	_, _ = collection6.DeleteMany(context.TODO(), bson.M{})
	_, _ = collection7.DeleteMany(context.TODO(), bson.M{})

}
