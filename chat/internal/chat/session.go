package chat

import (
	"github.com/google/uuid"
)

type Session struct {
	ID     uuid.UUID
	UserID string
	Conn   ConnGateway
}

func (s *Session) ReadPump() error {
	for {
		data, err := s.Conn.ReadConn()
		if err != nil {
			return err
		}
		switch data.Type {
		case Msg:
		default:
			return ErrorInvalidMessageType
		}
	}
}

func (s *Session) WritePump() {
}
