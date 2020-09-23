package HandleRequest

import (
	"Joker/query"
	"html/template"
	"net/http"
)
//随机查询
func Visit(w http.ResponseWriter,request *http.Request)  {

	query.Power()

	str := "random.html"
	temp, err := template.ParseFiles("./HTML/" + str)

	if err != nil {
		errorTemp, _ := template.ParseFiles("./HTML/error.html")
		errorTemp.Execute(w, err.Error())
		return
	}
	//fmt.Println(query.SuijiDatas)
	temp.Execute(w, query.SuijiDatas)

}
