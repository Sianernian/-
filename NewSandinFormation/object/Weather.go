package object

type Weather struct {
	Reason string //请求是否成功
	Result Result_weather //数据
	Error_code int //状态码
}

type Result_weather struct {
	City string //城市
	Realtime Realtime //实时(当时天气)
	Future []Future //未来(未来天气)
}

type Realtime struct {
	Temperature string //温度
	Humidity string //湿度
	Info string //信息
	Wid string //风
	Direct string //
	Power string //等级
	Aqi string
}

type Future struct {
	Date string //日期
	Temperature string //温度
	Weather string //天气清况
	Wid Wid //风力
	Direct string
}

type Wid struct {
	Day string //白天
	Night string //晚上
}