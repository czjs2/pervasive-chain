package ws

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

type BaseMsg struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
