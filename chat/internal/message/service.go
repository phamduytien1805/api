package message

import (
	"context"
	"encoding/json"
	"log/slog"
)

type MessageService interface {
	BroadcastTextMessage(ctx context.Context, payload MessagePayload) error
}

type MessageServiceImpl struct {
	logger    *slog.Logger
	publisher PubGateway
}

func NewMessageService(logger *slog.Logger, pub PubGateway) MessageService {
	return &MessageServiceImpl{
		logger:    logger,
		publisher: pub,
	}
}

func (s *MessageServiceImpl) BroadcastTextMessage(ctx context.Context, payload MessagePayload) error {
	textMsg, err := mapToTextMessage(payload)
	if err != nil {
		return err
	}
	rawMsg, err := json.Marshal(textMsg)
	if err != nil {
		return err
	}
	return s.publisher.PublishMessage(rawMsg)
}
