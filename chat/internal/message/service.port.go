package message

import (
	"context"
)

type MessageService interface {
	BroadcastMessage(ctx context.Context, payload *MessagePayload) error
}
