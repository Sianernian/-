package network

import (
	"NewSandinFormation/object"
	"fmt"
	"html/template"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {

	//1、执行网络请求
	_,err := http.Get("http:127.0.0.1:9092")


	//2、判断网络请求既读取网络请求数据
	/**	2.0 如果遇到错误
	*	parse一个eeror错误页面
	*	显示错误页面
	*	return
	*/
	if err == nil {
		fmt.Println(err.Error())
		return
	}else {
		//2.1 如果成功
		//准备模板文件html
		temp := template.Must(template.ParseFiles("./views/pages/newSandinFormation.html"))


		//fmt.Println("hello world")
		//除第一次以外，剩下的请求，都会阻塞在此处
		ob := object.OB{
			New: <- Ch_new,
			Weather: <- Ch_weather,
			Stroy: <- Ch_stroy,
		}
		//fmt.Println(ob.New.Result.Data[0].Thumbnail_pic_s)
		//把模板页面显示到前端浏览器
		temp.Execute(writer,ob)
	}

}
