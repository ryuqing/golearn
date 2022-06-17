package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
	"github.com/task/lib"
)



var mapMarkets map[string]config.MarketType
func init()  {
	mapMarkets, err = config.GetConfig("mxc", "_")
	if err != nil {
		log.Printf("get config error")
	}
}

var addr = flag.String("addr", "wbs.mexc.com", "http service address")
func main() {

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/raw/ws"}

	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial fail:", err)
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() { //开启goroutine
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second) //创建z周期为1s的定时器
	defer ticker.Stop()

	var depthSubStr string

	for _, market := range mapMarkets{
		depthSubStr = "{\"op\":\"sub.limit.depth\", \"depth\":5, \"symbol\":\"" + market.Coin + "_" + market.Currency + "\"}"
		fmt.Println(depthSubStr)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(depthSubStr)); err != nil {
			log.Println(err)
		}
	}

	for {
		select {
		case <-done:
			return
		}
	}
}


func dealDepthData(msg string)  {
	var rawData map[string]interface{}

	err = json.Unmarshal([]byte(msg), &rawData)
	if err != nil {
		log.Println(err)
		return
	}

	if _, hasSymbol := rawData["symbol"]; !hasSymbol {
		log.Println("no symbol")
		return
	}

	if _, hasChan := rawData["channel"]; !hasChan {
		log.Println("no data")
		return
	}

	if _, hasData := rawData["data"]; !hasData {
		log.Println("no data")
		return
	}

	channel := rawData["channel"].(string)
	if _, hasConfig := mapMarkets[channel]; !hasConfig {
		log.Println("no config")
		return
	}

	data := rawData["data"].(map[string]interface{})
	bids := data["bids"].([]interface{})
	asks := data["asks"].([]interface{})

}


func formatDepth(data interface{}) [][]string {
	var itemDepth = make([]string, 2)
	var responseList = make([][]string, 1);
	list := data.([]interface{})

	for _, originData := range list {
		itemData := originData.(map[string]interface{})
		itemDepth[0] = itemData["p"].(string)
		itemDepth[1] = itemData["q"].(string)
		responseList = append(responseList, itemDepth)
	}
	return responseList;
}



