package main

import (
	"NewSandinFormation/network"
	"fmt"
	"net/http"
)


func main() {
	//文件路径
	newUrl := "http://v.juhe.cn/toutiao/index?type=&key=c7d1ba895699743816ba183d865e72a6"
	weatherUrl := "http://apis.juhe.cn/simpleWeather/query?city=%E5%8D%97%E6%98%8C&key=8859f1da05728f44faf07031b0ea99c2"
	storyUrl := "http://v.juhe.cn/joke/randJoke.php?key=ed2bf241fb8b3e8dc6a745ab150c0abd"

	//apiname

	Apiurl :=  []string{newUrl,weatherUrl,storyUrl}

	//将所要获取的数据的url和其对应的struct名传入获取数据
	//3个协程
	for n := 0; n < len(Apiurl); n++ {
		go network.GetNetworkData(Apiurl,n)
	}

	//加载静态页面
	http.Handle("/views/img/", http.StripPrefix("/views/img/", http.FileServer(http.Dir("./views/img/"))))//加载图片
	http.Handle("/views/css/", http.StripPrefix("/views/css/", http.FileServer(http.Dir("./views/css/"))))//加载css样式表
	http.Handle("/views/js/", http.StripPrefix("/views/js/", http.FileServer(http.Dir("./views/js/"))))//加载JQ和JS

	//加载动态页面
	//http://127.0.0.1:9092/views/iamge/
	http.HandleFunc("/index",network.Index)

	//添加服务
	fmt.Println("进入新闻资讯")
	http.ListenAndServe(":9092",nil)
}