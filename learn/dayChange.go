package learn


/*实现查找连续上涨的币种*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type config struct {
	Code int `json:"code"`
	Data configData `json:"data"`
}

type configData struct {
	Markets map[int]market `json:"markets"` // 注意这里的字段要与返回的字段对应上！
}

type market struct {
	Site     string
	Coin     string
	Currency string
}
type ticker struct {
	Id int
	Last float32
	Open float32
	Time int16
}
type klineChange struct {
	Date string
	ChangeRate float64
}

func main() {
	//uploadJSON := `{"xxxx": "test","zzzz": "1111111"}`
	//var jsonSlice2 map[string]interface{}
	//json.Unmarshal([]byte(uploadJSON), &jsonSlice2)
	//fmt.Println(jsonSlice2)

	checkSite := "huobi.pro"
	config := getMarketConfig()
	for marektId,market := range config.Markets {
		if market.Site != checkSite {
			continue;
		}
		fmt.Println("start_cron" + strconv.Itoa(marektId))
		result := getKline(marektId)
		if len(result) > 0 {
			fmt.Printf("%+v\n", market.Coin + "-" + market.Currency)
			fmt.Println(result)
			fmt.Print("---------------------------")
		}
		time.Sleep(time.Duration(1)*time.Second)
	}

}


func getMarketConfig() configData {
	url := "https://x.szsing.com/v2/quote/price/m_config?version=0"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	bodyByte, err := ioutil.ReadAll(response.Body) //body数据流
	//bodyStr := string(bodyByte)

	var params config
	err = json.Unmarshal(bodyByte, &params)

	return params.Data
}

func getKline(marketId int) []klineChange {

	nowTime := time.Now().Unix()
	var dayBefore int64
	dayBefore = 5
	dayBeforeTime := strconv.FormatInt(nowTime - nowTime%(24*3600) - dayBefore * 24 * 3600,  10)
	url := "https://x.szsing.com/v2/quote/chart/m_candlestick?type=day&market_id=" + strconv.Itoa(marketId) + "&since=" + dayBeforeTime + "000"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	bodyByte, err := ioutil.ReadAll(response.Body) //body数据流
	if err != nil {
		fmt.Println(err)
	}
	var jsonSlice2 map[string]interface{} //go语言结构体可以接受任意类型数据,比较实用
	err = json.Unmarshal(bodyByte, &jsonSlice2)
	if err != nil {
		fmt.Println(err)
	}

	klineData := jsonSlice2["data"].(map[string]interface{})//转换 需要命名转换
	klineList := klineData["list"].([]interface{})

	resultData := make([]klineChange, dayBefore)

	i := 0
	var riseDays int64
	riseDays = 0 //符合条件的上涨天数

	for _, item := range klineList {
		dayKline := item.([]interface{})
		dayRiseRate := (dayKline[4].(float64) - dayKline[1].(float64)) / dayKline[1].(float64)
		resultItem := klineChange{
			Date: time.Unix(int64(dayKline[0].(float64)/1000), 0).Format("2006-01-02 15:04:05"),
			ChangeRate: dayRiseRate,
		}
		if dayRiseRate > 0.1 && dayRiseRate < 0.3 { //这里是范围
			riseDays ++
		}
		resultData[i] = resultItem
		i++

	}

	if riseDays >= dayBefore - 2 {
		return resultData
	} else  {
		return nil
	}
}




/**
******** 这是小记 ************



fmt.Printf("%+v\n", params) //这个可以打印完整对象 （https://blog.cyeam.com/golang/2017/03/06/go-fmt-v）--这个比较实用
JSON处理示例：查看https://www.cnblogs.com/Detector/p/9048678.html; https://www.cnblogs.com/liuhe688/p/10971332.html; 这里作学习实用
如果想快速拿到json里的数据可以用这个包

1.interface 可以承载任意类型的数据
2.string(marketId) string()会直接把字节或者数字转换为字符的UTF-8表现形式
*/
 