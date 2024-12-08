package hub

import "log/slog"

type handleOnMessage func()
type handleOnConnect func()
type handleOnDisconnect func()

type Hub struct {
	logger       *slog.Logger
	onConnect    handleOnConnect
	onMessage    handleOnMessage
	onDisconnect handleOnDisconnect
}

func NewHub(logger *slog.Logger) *Hub {
	return &Hub{
		logger:       logger,
		onConnect:    func() {},
		onMessage:    func() {},
		onDisconnect: func() {},
	}
}

func (h *Hub) OnConnect(fn func()) {
	h.onConnect = fn
}

func (h *Hub) OnMessage(fn func()) {
	h.onMessage = fn
}

func (h *Hub) OnDisconnect(fn func()) {
	h.onDisconnect = fn
}
