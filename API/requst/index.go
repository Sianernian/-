package requst

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {


	_,err:=http.Get("http:127.0.0.1:9000")

	if err ==nil{
		fmt.Println(err.Error())

		return
	}else{
		tempt, _ := template.ParseFiles("./view/HD.html")

		a :=<-Ch

		tempt.Execute(writer,a)
	}

}
