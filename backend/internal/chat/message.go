package chat

type Message struct {
	FromUID int64  `json:"from_uid"`
	ToUID   int64  `json:"to_uid"`
	Body    string `json:"body"`
	Type    string `json:"type,omitempty"`
	Typing  bool   `json:"typing,omitempty"`
}
