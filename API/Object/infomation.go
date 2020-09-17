package Object

type Infomation struct {
	Id string	`json:"_id"`
	Title string	`json:"title"`
	Year int	`json:"year"`
	Month int	`json:"month"`
	Day int		`json:"day"`
	Des string	`json:"des"`
	Lunar string `json:"lunar"`
	Pic string 	`json:"pic"`
}
