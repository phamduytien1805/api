package chat

import "github.com/phamduytien1805/chatmodule/internal/message"

type ConnGateway interface {
	WriteConn(data interface{}) error
	ReadConn() (message.BaseEvent, error)
	HandleError(error)
}
