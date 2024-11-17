package chat

import "errors"

var (
	ErrorInvalidMessage             = errors.New("Message is invalid")
	ErrorInvalidMessageType         = errors.New("Message type is invalid")
	ErrorInitializeSession          = errors.New("Cannot initialize session")
	ErrorInitializeReader           = errors.New("Cannot initialize reader")
	ErrorHandleMessage              = errors.New("Cannot handle message")
	ErrorHandleBroadcastTextMessage = errors.New("Cannot handle to broadcast text message")
)
