package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"tools"
)

/**
这个程序是检查币安的公告 有新的公告发邮件
 */
func main() {
	for true {
		run()
		time.Sleep(time.Duration(15)*time.Second)
	}
}

func run()  {
	res, err := http.Get("https://www.binancezh.sh/zh-CN/support/announcement/c-48?navId=48")
	if err != nil || res.StatusCode != 200  {
		log.Fatal("")
	}

	//判断并创建文件
	fileName := "./records.txt"
	_, err = os.Stat(fileName)
	if err != nil && os.IsNotExist(err) {
		f,err := os.Create(fileName)
		defer f.Close()
		if err != nil {
			log.Fatalln("文件创建失败", err)
		}
	}

	//读取
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("读取记录失败", err);
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)


	var existRecords []string
	err = json.Unmarshal(content, &existRecords)


	firstCont,contentLink := getBanFirstAnnounce();
	if len(firstCont) < 1 {
		log.Fatalln("获取币安信息失败")
	}


	exitRecordLength := len(existRecords);
	if exitRecordLength == 0 || existRecords[exitRecordLength -1] != firstCont {
		existRecords = append(existRecords, firstCont)

		stringTo, err := json.Marshal(existRecords);
		err = ioutil.WriteFile(fileName, stringTo, 0666)
		if err != nil {
			fmt.Println("ioutil WriteFile error: ", stringTo)
		}
		sendCloud.SendMail(firstCont, contentLink)
	}
}

func getBanFirstAnnounce()(string, string) {
	res, err := http.Get("https://www.binancezh.sh/zh-CN/support/announcement/c-48?navId=48")
	if err != nil || res.StatusCode != 200  {
		log.Fatal("")
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("no content return")
	}
	firstCont := doc.Find(".css-1wr4jig .css-6f91y1 .css-1ej4hfo").First().Text()
	val, exists := doc.Find(".css-1wr4jig .css-6f91y1 .css-1ej4hfo").First().Attr("href")
	if !exists {
		fmt.Println("no data")
	}
	val = "https://www.binancezh.sh/" + val;
	return firstCont,val;
}
