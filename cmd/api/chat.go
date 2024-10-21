package api

import (
	"net/http"

	"github.com/coder/websocket"
	"github.com/phamduytien1805/pkg/config"
)

type WSConn struct {
	*websocket.Conn
	config *config.Config
}

func WebSocketBuilder(config *config.Config) WSConn {
	return WSConn{
		config: config,
	}
}

func (app *application) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	defer conn.CloseNow()

}
