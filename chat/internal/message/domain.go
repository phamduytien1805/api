package message

import "github.com/phamduytien1805/pkgmodule/id_generator"

type TextMessage struct {
	RoomID    id_generator.UUID `json:"room_id"`
	From      id_generator.UUID `json:"from"`
	Text      string            `json:"text"`
	Timestamp int64             `json:"timestamp"`
}
