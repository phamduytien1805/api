package chat

import "errors"

var (
	ErrorInvalidMessageType = errors.New("Message type is invalid")
	ErrorInitializeSession  = errors.New("Cannot initialize session")
	ErrorHandleMessage      = errors.New("Cannot handle message")
)
