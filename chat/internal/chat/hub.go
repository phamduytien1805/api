package chat

type Hub struct {
	sessions map[*Session]bool
	rooms    map[string]*Room
}

func NewHub() *Hub {
	return &Hub{
		sessions: make(map[*Session]bool),
		rooms:    make(map[string]*Room),
	}
}

func (h *Hub) Run() {
	for {
		select {}
	}
}

func (h *Hub) OnConnect(session *Session) {
	h.sessions[session] = true
}
