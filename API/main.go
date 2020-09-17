package main

import (
	"API/requst"
	"fmt"
	"net/http"
)

func main() {

	// 复杂模式

	//err , Page :=requst.Page("http://api.juheapi.com/japi/toh?key=b761769ce503167b0bd492ab6e5ccce0&v=1.0&month=11&day=1")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	// 简单 模式
	err := requst.Page()
	if err != nil {
		fmt.Println(err.Error())
	}

	//服务接口
    http.HandleFunc("/index", requst.Index)

	http.Handle("/view/js/",http.StripPrefix("/view/js/",http.FileServer(http.Dir("./view/js/"))))

    err = http.ListenAndServe(":9000" , nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Over")
}

