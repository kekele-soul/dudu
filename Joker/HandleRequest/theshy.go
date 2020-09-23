package HandleRequest

import (
	"Joker/query"
	"html/template"
	"net/http"
)
//最新笑话
func Shy(w http.ResponseWriter,request *http.Request) {

	query.Newest()

	str :="new.html"
	tepm,err :=template.ParseFiles("./HTML/"+str)
	if err != nil {
		errTemp,_:=	template.ParseFiles("./HTML/err.html")
		errTemp.Execute(w,err.Error())
		return
	}
	tepm.Execute(w,query.NewDatas)
}
