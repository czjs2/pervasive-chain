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
	go socketClient(interrupt)
	select {}
}

func apiClient(interrupt chan os.Signal) {

	cleanDbData()

	bHeartform := form.HeartBeatFrom{
		Type:   "b",
		Number: "10000",
		Id:     "10000",
		Time:   utils.GetNowTime(),
	}

	rHeartform := form.HeartBeatFrom{
		Type:   "r",
		Number: "20000",
		Id:     "20000",
		Time:   utils.GetNowTime(),
	}

	sHeartform := form.HeartBeatFrom{
		Type:   "s",
		Number: "30000",
		Id:     "30000",
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
		Number:   "30000",
		Id:       "30000",
		Height:   100,
		Father:   "99",
		Hash:     "100",
		Vrf:      "100",
		Time:     utils.GetNowTime(),
		Interval: 100,
		Trans:    100,
		Size:     100,
		Detail:   nil,
	}

	_, err = ReportBlock(host, "/v1.0/block", "", reportBlockForm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	reportFlowForm := form.ReportFlowForm{
		Type:   "s",
		Number: "30000",
		Id:     "30000",
		Time:   utils.GetNowTime(),
		In:     100,
		Out:    100,
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
	defer c.Close()
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {

			select {
			case <-interrupt:
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
	ticker := time.NewTicker(2 * time.Second)

	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			res := `{"uri":"chainInfo","body":{},"msgId":"msgid111111111"}`
			err := c.WriteMessage(websocket.TextMessage, []byte(res))
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
