package HandleRequest

import (
	"Joker/query"
	"html/template"
	"net/http"
)

//按时间查询
func Xfei(w http.ResponseWriter,request *http.Request) {

	query.Bytie()

	str :="login.html"
	tepm,err :=template.ParseFiles("./HTML/"+str)
	if err != nil {
	errTemp,_:=	template.ParseFiles("./HTML/err.html")
	errTemp.Execute(w,err.Error())
		return
	}
	//fmt.Println(query.Date)
	tepm.Execute(w, query.TimeDatas)
}
