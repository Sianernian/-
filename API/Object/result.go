package Object


type Result  struct {
    Info  []Infomation  `json:"result"`
    Reason string   `json:"reason"`
    ErrorCode int `json:"error_code"`

}