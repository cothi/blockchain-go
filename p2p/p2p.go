package p2p

import (
	"net/http"
	"tetgo/tetgocoin/utill"

	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn
var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	conns = append(conns, conn)
	utill.HandleErr(err)

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}

		for _, aConn := range conns {
			if aConn != conn {
				utill.HandleErr(aConn.WriteMessage(websocket.TextMessage, p))
			}
		}
	}
}
