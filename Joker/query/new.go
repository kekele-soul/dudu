package query

import (
	"Joker/strustt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var NewDatas = strustt.To{}

//最新笑话
func Newest() {
	client := http.Client{}
	//url :="http://'v.'juhe.cn/joke/content/list.php?key=e622e0a6104326bc52c759b8661a3a61&page=2&pagesize=10&sort=asc&time=1418745237"
	url :="http://v.juhe.cn/joke/content/text.php?page=1&pagesize=20&key=b474e38a314c9f468664955665fa6815"
	request, _ := http.NewRequest("GET", url, nil)
	response, _ := client.Do(request)

	//fmt.Println(response.Body)
	Bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("最新出错:",err.Error())
		return
	}

	//fmt.Println(string(Bytes))
	json.Unmarshal(Bytes, &NewDatas)
	fmt.Println("最新",NewDatas)

}
