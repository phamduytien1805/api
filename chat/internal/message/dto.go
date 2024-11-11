package message

type BaseEvent struct {
	Text *MessagePayload `json:"text,omitempty"`
	Join *JoinPayload    `json:"join,omitempty"`
}

type MessagePayload struct {
	Text   string `json:"text"`
	RoomID string `json:"room_id"`
	From   string `json:"from"`
}

type JoinPayload struct {
	RoomID string `json:"room_id"`
	From   string `json:"from"`
}
