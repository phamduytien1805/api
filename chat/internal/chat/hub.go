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

func NewHub(logger *slog.Logger, msgSvc message.MessageService) *Hub {
	return &Hub{
		logger:   logger,
		sessions: make(map[*Session]bool),
		rooms:    make(map[string]*Room),
		msgSvc:   msgSvc,
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
		conn.HandleError(ErrorInitializeSession)
		return
	}
	session := &Session{
		ID:     sessionId,
		Conn:   conn,
		UserID: "demo",
		Hub:    h,
	}

	h.logger.Info("New session", "sessionId", sessionId)

	go session.WritePump()
	err = session.ReadPump()
	if err != nil {
		h.logger.Error("Cannot run read pump", "detailed", err)
		conn.HandleError(ErrorInitializeReader)
		return
	}
}

func (h *Hub) onConnect(session *Session) {
	h.sessions[session] = true
}

func (h *Hub) onMessage(session *Session, data []byte) {
	eventMessage, err := mapRawToBaseEvent(data)
	if err != nil {
		h.logger.Error("Cannot map raw data to event message", "detailed", err)
		session.Conn.HandleError(ErrorInvalidMessage)
		return
	}
	switch {
	case eventMessage.Text != nil:
		if err := h.msgSvc.BroadcastTextMessage(context.Background(), *eventMessage.Text); err != nil {
			h.logger.Error("Cannot broadcast text message", "detailed", err)
			session.Conn.HandleError(ErrorHandleBroadcastTextMessage)
			return
		}
	default:
	}
}
