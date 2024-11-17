package message

import (
	"time"

	"github.com/phamduytien1805/pkgmodule/id_generator"
)

func mapToTextMessage(payload MessagePayload) (TextMessage, error) {
	roomId, err := id_generator.Parse(payload.RoomID)
	if err != nil {
		return TextMessage{}, err
	}
	from, err := id_generator.Parse(payload.From)
	if err != nil {
		return TextMessage{}, err
	}
	return TextMessage{
		RoomID:    roomId,
		From:      from,
		Text:      payload.Text,
		Timestamp: time.Now().Unix(),
	}, nil
}
