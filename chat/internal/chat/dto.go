package chat

type BaseMessage struct {
	Type      MessageType `json:"type"`
	Timestamp int64       `json:"timestamp"`
}

type TextMessage struct {
	BaseMessage
	Text   string `json:"text"`
	RoomID string `json:"room_id"`
	From   string `json:"from"`
}

type JoinMessage struct {
	BaseMessage
	RoomID string `json:"room_id"`
	From   string `json:"from"`
}
