package query

import (
	"Joker/strustt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var TimeDatas = strustt.To{}
//按照更新时间查询
func Bytie() {
	client := http.Client{}
	url := "http://v.juhe.cn/joke/content/list.php?sort=&page=&pagesize=&time=1418816972&key=b474e38a314c9f468664955665fa6815"
	request, _ := http.NewRequest("GET", url, nil)
	response, _ := client.Do(request)
	Bytes, _ := ioutil.ReadAll(response.Body)
//	fmt.Println(string(Bytes))


	json.Unmarshal(Bytes, &TimeDatas)

	fmt.Println(TimeDatas.Result.Data[1].Content)
	fmt.Println("时间",TimeDatas)

	}

