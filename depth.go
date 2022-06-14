package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    var respData map[string]interface{}
    var marketConfigUrl string = "http://localhost:10085/v2/market/config?sites=bittrex"
    response, err := http.Get(marketConfigUrl)
    if err != nil {
        fmt.Println(err)
    }

    bodyByte, err := ioutil.ReadAll(response.Body) //body数据流
    if err != nil {
        fmt.Println(err)
    }

    err = json.Unmarshal(bodyByte, &respData)
    if err != nil {
        fmt.Println(err)
    }
    respData = respData["data"].(map[string]interface{})
    bitMarkets := respData["bittrex"].([]interface{})
    fmt.Printf("%+v\n", bitMarkets)
    for _,item := range bitMarkets{
        market := item.(map[string]interface{}) //类型转换
        fmt.Printf("%+v\n", market["id"])
    }

}