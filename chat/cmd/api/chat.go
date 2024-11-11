package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/coder/websocket"
	"github.com/phamduytien1805/chatmodule/internal/chat"
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

	app.hub.OnJoinHub(c)

	conn.Close(websocket.StatusNormalClosure, "connection closed")
}

type Conn struct {
	conn *websocket.Conn
}

func NewConn(conn *websocket.Conn) Conn {
	return Conn{
		conn: conn,
	}
}

func (c Conn) ReadConn() ([]byte, error) {
	msgType, data, err := c.conn.Read(context.Background())
	if err != nil {
		return nil, err
	}
	switch msgType {
	case websocket.MessageText:
		return data, nil
	case websocket.MessageBinary:
		// Ignore binary message
		return nil, nil
	}
	return nil, nil
}

func (c Conn) HandleError(err error) {
	if err != nil {
		c.conn.Close(websocket.StatusNormalClosure, "connection closed")
	}
	switch {
	case errors.Is(err, chat.ErrorHandleMessage) || errors.Is(err, chat.ErrorInvalidMessageType):
		c.conn.Close(websocket.StatusInvalidFramePayloadData, "invalid message")
	case errors.Is(err, chat.ErrorInitializeSession):
		c.conn.Close(websocket.StatusUnsupportedData, "fail to create session")
	default:
		c.conn.Close(websocket.StatusInternalError, "internal error")
	}
}
