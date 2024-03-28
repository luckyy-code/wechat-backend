package domain

import "encoding/json"

// 消息类型
const (
	USER_TYPE  = 1
	GROUP_TYPE = 2
)

// Message 消息格式
type Message struct {
	Content  string `json:"content"`
	From     int64  `json:"from"`
	To       int64  `json:"to"`
	Type     int64  `json:"type"` //群消息还是用户消息
	SendTime string `json:"send_time"`
}

func (m *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
