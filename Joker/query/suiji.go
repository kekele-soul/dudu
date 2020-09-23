package query

import (
	"Joker/strustt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var SuijiDatas = strustt.Randome{}
//随机查询
func Power()  {

	client :=http.Client{}
	//url :="http://v.juhe.cn/joke/content/list.php?key=e622e0a6104326bc52c759b8661a3a61&page=2&pagesize=10&sort=asc&time=1418745237"
	url :="http://v.juhe.cn/joke/randJoke.php?key=b474e38a314c9f468664955665fa6815"
	request,_:=http.NewRequest("GET",url,nil)
	response,_:=client.Do(request)
	Bytes ,_:=ioutil.ReadAll(response.Body)
	//fmt.Println(string(core))

	//fmt.Println(Data)
	json.Unmarshal(Bytes,&SuijiDatas)
	fmt.Println("随机",SuijiDatas)

}
