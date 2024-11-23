package chat

import (
	"github.com/phamduytien1805/pkgmodule/id_generator"
)

type ReplyMsg struct {
	ReplyTo string `json:"reply_to"`
	Ok      bool   `json:"ok"`
	Msg     string `json:"msg,omitempty"`
}

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

func (s *Session) Ok(eventId string) error {
	return s.Conn.WriteConn(ReplyMsg{
		ReplyTo: eventId,
		Ok:      true,
	})
}

func (s *Session) Error(eventId string, err error) error {
	return s.Conn.WriteConn(ReplyMsg{
		ReplyTo: eventId,
		Ok:      false,
		Msg:     err.Error(),
	})
}
