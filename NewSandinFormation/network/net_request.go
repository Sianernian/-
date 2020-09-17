package network

import (
	"NewSandinFormation/object"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var Ch_new chan object.New
var Ch_weather chan object.Weather
var Ch_stroy chan object.Stroy

func GetNetworkData (apiUrl []string,n int) (error) {
	//1.创建客户端
	client := http.Client{}

	request, err := http.NewRequest("GET",apiUrl[n],nil)
	if err !=nil {
		fmt.Println(err.Error())
		return  err
	}

	//开启request
	response, err := client.Do(request)
	if err !=nil {
		fmt.Println(err.Error())
		return  err
	}

	//获取数据体
	data_new, err := ioutil.ReadAll(response.Body)
	if err !=nil {
		fmt.Println(err.Error())
		return  err
	}

	switch n {
		case 0:
			Ch_new = make(chan object.New)
			var new object.New;
			//解析json数据
			err = json.Unmarshal(data_new,&new)
			if err !=nil {
				fmt.Println(err.Error())
				return  err
			}
			Ch_new <- new
		case 1:
			Ch_weather = make(chan object.Weather)
			var weather object.Weather;
			//解析json数据
			err = json.Unmarshal(data_new,&weather)
			if err !=nil {
				fmt.Println(err.Error())
				return  err
			}
			Ch_weather <- weather
		case 2:
			Ch_stroy = make(chan object.Stroy)
			var stroy object.Stroy;
			//解析json数据
			err = json.Unmarshal(data_new,&stroy)
			if err !=nil {
				fmt.Println(err.Error())
				return  err
			}
			Ch_stroy <- stroy
	}

	return nil
}



//将所要获取的数据的url和其对应的struct名传入获取数据
//func GetNetworkData_New (url string) (error) {
//	//1.创建客户端
//	client := http.Client{}
//
//	request, err := http.NewRequest("GET",url,nil)
//	if err !=nil {
//		fmt.Println(err.Error())
//		return  err
//	}
//
//	//开启request
//	response, err := client.Do(request)
//	if err !=nil {
//		fmt.Println(err.Error())
//		return  err
//	}
//
//	//获取数据体
//	data_new, err := ioutil.ReadAll(response.Body)
//	if err !=nil {
//		fmt.Println(err.Error())
//		return  err
//	}
//
//
//	var new object.New
//
//	//解析json数据
//	err = json.Unmarshal(data_new,&new)
//	if err !=nil {
//		fmt.Println(err.Error())
//		return  err
//	}
//	return nil
//}
//
//func GetNetworkData_Weather (url string) (error) {
//	//1.创建客户端
//	client := http.Client{}
//
//	request, err := http.NewRequest("GET",url,nil)
//	if err !=nil {
//		fmt.Println(err.Error())
//		return  err
//	}
//
//	//开启request
//	response, err := client.Do(request)
//	if err !=nil {
//		fmt.Println(err.Error())
//		return err
//	}
//
//	//获取数据体
//	data_weather, err := ioutil.ReadAll(response.Body)
//	if err !=nil {
//		fmt.Println(err.Error())
//		return err
//	}
//
//	var weather object.Weather
//
//	//解析json数据
//	err = json.Unmarshal(data_weather,&weather)
//	if err !=nil {
//		fmt.Println(err.Error())
//		return err
//	}
//	return nil
//}
//
//func GetNetworkData_Stroy (url string) (error) {
//	//1.创建客户端
//	client := http.Client{}
//
//	request, err := http.NewRequest("GET",url,nil)
//	if err !=nil {
//		fmt.Println(err.Error())
//		return  err
//	}
//
//	//开启request
//	response, err := client.Do(request)
//	if err !=nil {
//		fmt.Println(err.Error())
//		return err
//	}
//
//	//获取数据体
//	data_stroy, err := ioutil.ReadAll(response.Body)
//	if err !=nil {
//		fmt.Println(err.Error())
//		return  err
//	}
//
//	var stroy object.Stroy
//
//	//解析json数据
//	err = json.Unmarshal(data_stroy,&stroy)
//	if err !=nil {
//		fmt.Println(err.Error())
//		return  err
//	}
//	return nil
//}
