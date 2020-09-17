package requst

import (
	"API/Object"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var Ch chan Object.Result


//  复杂

//func Page(url string) (error, Object.Result){
//
//}


// 简单模式
func Page()(error){

	client := http.Client{}

	//METHOD 必须大写
	request, err := http.NewRequest("GET", "http://api.juheapi.com/japi/toh?key=b761769ce503167b0bd492ab6e5ccce0&v=1.0&month=11&day=1", nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	//resp, err := http.Get(url)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}

	Data, err := ioutil.ReadAll(resp.Body)

	var result Object.Result

	// json 转换
	err = json.Unmarshal(Data, &result)
	if err != nil {
		fmt.Println(err.Error())
	}
	Ch = make(chan Object.Result, 100000)

	Ch <- result
	return err
}

