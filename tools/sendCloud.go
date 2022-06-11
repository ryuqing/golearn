package sendCloud

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "bytes"
    "net/url"
    //"strings"
)



func SendMail(subject string, html string){
	RequestURI := "http://api.sendcloud.net/apiv2/mail/send"
	PostParams := url.Values{
		"apiUser": {"513826081_test_yEAgnH"},
		"apiKey":  {"d3b41e19b6ca794c51d5ebf7aa4054e8"},
		"from":     {"operationcat@qq.com"},
		"fromName": {"机器人go"},
		"to":       {"operationcat@qq.com"},
		"subject":  {subject},
		"html":     {html},
	}
	PostBody := bytes.NewBufferString(PostParams.Encode())
	ResponseHandler, err := http.Post(RequestURI,"application/x-www-form-urlencoded",PostBody)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ResponseHandler.Body.Close()
	BodyByte, err := ioutil.ReadAll(ResponseHandler.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(BodyByte))
}
