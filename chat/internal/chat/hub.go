package chat

import (
	"context"
	"log/slog"

	"github.com/phamduytien1805/chatmodule/internal/message"
	"github.com/phamduytien1805/pkgmodule/id_generator"
)

type Hub struct {
	logger   *slog.Logger
	sessions map[*Session]bool
	rooms    map[string]*Room
	msgSvc   message.MessageService
}

func NewHub(logger *slog.Logger) *Hub {
	return &Hub{
		logger:   logger,
		sessions: make(map[*Session]bool),
		rooms:    make(map[string]*Room),
	}
}

func (h *Hub) Run() {
	for {
		select {}
	}
}

func (h *Hub) OnJoinHub(conn ConnGateway) {
	sessionId, err := id_generator.NewUUID()
	if err != nil {
		h.logger.Error("Cannot create sessionId", "detailed", err)
		conn.HandleError(err)
		return
	}
	session := &Session{
		ID:     sessionId,
		Conn:   conn,
		UserID: "demo",
		Hub:    h,
	}

	go session.WritePump()
	err = session.ReadPump()
	conn.HandleError(err)
}

func (h *Hub) onConnect(session *Session) {
	h.sessions[session] = true
}

func (h *Hub) onMessage(session *Session, data []byte) {
	eventMessage, err := mapRawToBaseEvent(data)
	if err != nil {
		session.Conn.HandleError(err)
		return
	}
	switch {
	case eventMessage.Text != nil:
		h.msgSvc.BroadcastMessage(context.Background(), eventMessage.Text)
	default:
	}
}
