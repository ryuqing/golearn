package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MarketType struct {
	Id string
	Site string
	Coin string
	Currency string
}

var mapMarkets = make(map[string]MarketType)

func GetConfig(site string, character string) (map[string]MarketType, error) {
	var respData map[string]interface{}
	var marketConfigUrl string = "http://api-master/v2/market/config?sites=" + site
	response, err := http.Get(marketConfigUrl)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return mapMarkets, err
	}

	bodyByte, err := ioutil.ReadAll(response.Body) //body数据流
	if err != nil {
		log.Println(err)
		return mapMarkets, err
	}

	err = json.Unmarshal(bodyByte, &respData)
	if err != nil {
		log.Println(err)
		return mapMarkets, err
	}

	if respData["code"] != 0.0 { //respData["code"] 变成了float32
		log.Println("response error", respData)
		return mapMarkets, err
	}

	respData = respData["data"].(map[string]interface{})
	siteConfig := respData[site].([]interface{})

	var market MarketType
	for _, item := range siteConfig {
		config := item.(map[string]interface{})
		pair := config["coin"].(string) + character + config["currency"].(string)
		market.Id = config["id"].(string)
		market.Coin = config["coin"].(string)
		market.Currency = config["currency"].(string)
		market.Site = site

		mapMarkets[pair] = market
	}

	return mapMarkets, nil
}


