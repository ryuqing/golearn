package main

import (
	"encoding/json"
	"fmt"
	"github.com/task/lib"
	"log"
)

func main()  {
	var str = "{\"bool\":false}"
	var a map[string]interface{}

	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println(err)
	}

	if _,ok := a["bool"]; ok {
		println(ok)
		println(11111)
	} else  {
		println(ok)
		log.Println(2222)
	}
	testConfig()
}


func testConfig()  {
	configData, err := config.GetConfig("kraken", "-")
	if (err != nil) {
		fmt.Println("there have an error")
		return
	}

	for key, item := range configData {
		fmt.Println(key)
		fmt.Printf("%+v\n", item.Coin)
	}
}