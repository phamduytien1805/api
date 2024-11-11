package chat

import (
	"github.com/google/uuid"
)

type Session struct {
	ID     uuid.UUID
	UserID string
	Conn   ConnGateway
	Hub    *Hub
}

func (s *Session) ReadPump() error {
	for {
		data, err := s.Conn.ReadConn()
		if err != nil {
			return err
		}
		s.Hub.onMessage(s, data)
	}
}

func (s *Session) WritePump() {
}
