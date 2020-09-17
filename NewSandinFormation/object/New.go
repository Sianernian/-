package object


type New struct {
	Reason string `json:"reason"` //返回信息是否成功
	Result Result_new //返回的数据的结果
	Error_code int //错误代码
}

type Result_new struct {
	Stat string //状态码
	Data []Data_new //数据
}

type Data_new struct {
	Uniquekey string //唯一密钥
	Title string //标题
	Date string //日期
	Category string //类别
	Author_name string //作者姓名
	Url string //网址
	Thumbnail_pic_s string //缩略图
	Thumbnail_pic_s02 string
	Thumbnail_pic_s03 string
}


