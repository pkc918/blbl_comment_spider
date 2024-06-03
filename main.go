package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Header struct {
	UserAgent string
}

func main() {
	url := "https://api.bilibili.com/x/v2/reply/wbi/main?oid=420981979&type=1&mode=3&pagination_str=%7B%22offset%22:%22%22%7D&plat=1&seek_rpid=&web_location=1315875&w_rid=aa722bbb349ba2ba3c3100f46d12aee0&wts=1717418420"
	header := Header{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36",
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", header.UserAgent)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Http get error is ", err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal("Http status code is ", res.StatusCode)
	}
	//fmt.Printf("%v", res.Body)
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	var jsonData map[string]interface{} // 定义一个 map 来存储 JSON 数据
	err = json.Unmarshal(bodyBytes, &jsonData)
	if err != nil {
		log.Fatal("Error unmarshalling JSON data: ", err)
	}

	if data, ok := jsonData["data"].(map[string]interface{}); ok {
		if replies, repliesOk := data["replies"].([]interface{}); repliesOk {
			for i, v := range replies {
				if reply, rok := v.(map[string]interface{}); rok {
					content, _ := reply["content"].(map[string]interface{})
					for key, value := range content {
						if key == "message" {
							fmt.Printf("第%v条评论：%v\n", i, value)
						}
					}
				}
			}
		}
	}
}
