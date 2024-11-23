package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/phamduytien1805/chatmodule/internal/chat"
	"github.com/phamduytien1805/chatmodule/internal/message"
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

func (c Conn) ReadConn() (message.BaseEvent, error) {
	var v message.BaseEvent
	err := wsjson.Read(context.Background(), c.conn, &v)
	if err != nil {
		return v, err
	}
	return v, nil
}

func (c Conn) WriteConn(data interface{}) error {
	return wsjson.Write(context.Background(), c.conn, data)
}

func (c Conn) HandleError(err error) {
	switch {
	case errors.Is(err, chat.ErrorHandleMessage) || errors.Is(err, chat.ErrorInvalidMessageType):
		c.conn.Close(websocket.StatusInvalidFramePayloadData, err.Error())
	case errors.Is(err, chat.ErrorInitializeSession):
		c.conn.Close(websocket.StatusUnsupportedData, err.Error())
	default:
		c.conn.Close(websocket.StatusInternalError, err.Error())
	}
}
