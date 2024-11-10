package api

import (
	"context"
	"net/http"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/phamduytien1805/chatmodule/internal/chat"
	"github.com/phamduytien1805/pkgmodule/id_generator"
)

func (app *application) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		app.logError(r, err)
		app.badRequestResponse(w, r, err)
		return
	}
	defer conn.CloseNow()

	c := NewConn(conn)

	sessionId, err := id_generator.NewUUID()
	if err != nil {
		app.logError(r, err)
		conn.Close(websocket.StatusInternalError, "Cannot create session")
		return
	}

	session := &chat.Session{
		ID:     sessionId,
		Conn:   c,
		UserID: "demo",
	}

	app.hub.OnConnect(session)

	go session.WritePump()
	err = session.ReadPump()
	if err != nil {
		app.logError(r, err)
		conn.Close(websocket.StatusInternalError, "Cannot read message from session")
		return
	}

	conn.Close(websocket.StatusNormalClosure, "")

}

type Conn struct {
	conn *websocket.Conn
}

func NewConn(conn *websocket.Conn) Conn {
	return Conn{
		conn: conn,
	}
}

func (c Conn) ReadConn() (interface{}, error) {
	var v interface{}
	err := wsjson.Read(context.Background(), c.conn, &v)
	if err != nil {
		return v, err
	}
	return v, nil
}
