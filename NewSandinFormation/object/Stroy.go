package object

type Stroy struct {
	Reason string //请求是否成功
	Result []Result_stroy //数据
	Error_code int //状态码
}

type Result_stroy struct {
	Content string //内容
	HashId string //哈希Id
	Unixtime int //时间戳
}