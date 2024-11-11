package chat

import (
	"encoding/json"

	"github.com/phamduytien1805/chatmodule/internal/message"
)

func mapRawToBaseEvent(data []byte) (*message.BaseEvent, error) {
	var msg message.BaseEvent
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
