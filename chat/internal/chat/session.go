package chat

import (
	"github.com/phamduytien1805/pkgmodule/id_generator"
)

type Session struct {
	ID     id_generator.UUID
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
