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

var addr = flag.String("addr", "127.0.0.1:9999", "http service address")
var host = "http://127.0.0.1:9999"

func main() {
	flag.Parse()
	interrupt := make(chan os.Signal, 1)
	interrupt1 := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt1, os.Interrupt)
	go apiClient(interrupt1)
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

var height int64 = 1

func reportData() {
	bHeartform := form.HeartBeatFrom{
		Type:   "b",
		Number: fmt.Sprintf("10000%d", utils.Rand(5)),
		Id:     fmt.Sprintf("10000%d", utils.Rand(5)),
		Time:   utils.GetNowTime(),
	}
	rHeartform := form.HeartBeatFrom{
		Type:   "r",
		Number: fmt.Sprintf("10000%d", utils.Rand(5)),
		Id:     fmt.Sprintf("10000%d", utils.Rand(5)),
		Time:   utils.GetNowTime(),
	}
	sHeartform := form.HeartBeatFrom{
		Type:   "s",
		Number: fmt.Sprintf("10000%d", utils.Rand(5)),
		Id:     fmt.Sprintf("10000%d", utils.Rand(5)),
		Time:   utils.GetNowTime(),
	}
	_, err := HeartBeat(host, "/v1.0/headbeat", "", bHeartform)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = HeartBeat(host, "/v1.0/headbeat", "", rHeartform)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = HeartBeat(host, "/v1.0/headbeat", "", sHeartform)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reportBlockForm := form.ReportBlockForm{
		Type:     "s",
		Number:   fmt.Sprintf("10000%d", utils.Rand(5)),
		Id:       fmt.Sprintf("10000%d", utils.Rand(5)),
		Height:   height,
		Father:   "fatherHash",
		Hash:     "hash",
		Vrf:      "vrf",
		Time:     utils.GetNowTime(),
		Interval: utils.Rand(100),
		Trans:    utils.Rand(10000),
		Size:     utils.Rand(10000),
		Detail:   nil,
	}
	height = height+1
	_, err = ReportBlock(host, "/v1.0/block", "", reportBlockForm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reportFlowForm := form.ReportFlowForm{
		Type:   "s",
		Number: fmt.Sprintf("10000%d", utils.Rand(5)),
		Id:     fmt.Sprintf("10000%d", utils.Rand(5)),
		Time:   utils.GetNowTime(),
		In:     utils.Rand(1000),
		Out:    utils.Rand(1000),
	}
	_, err = ReportFlow(host, "/v1.0/flow", "", reportFlowForm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
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

			chainInfoCmd := fmt.Sprintf(`{"uri":"chainInfo","body":{},"msgId":"msgId%d"}`, time.Now().Nanosecond())
			err = c.WriteMessage(websocket.TextMessage, []byte(chainInfoCmd))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			blockInfoCmd := fmt.Sprintf(`{"uri":"blockInfo","body":{},"msgId":"msgId%d"}`, time.Now().Nanosecond())
			err = c.WriteMessage(websocket.TextMessage, []byte(blockInfoCmd))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			cmdInfo := fmt.Sprintf(`{"uri":"cmd","body":{"key":{"trans":1000}},"msgId":"msgId%d"}`, time.Now().Nanosecond())
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
	collection1 := db.Collection(db.ChainInfoTable)
	collection2 := db.Collection(db.TotalChainTable)
	collection3 := db.Collection(db.HistoryChainInfoTable)
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
