package p2p

import (
	"net/http"
	"tetgo/tetgocoin/utill"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	_, err := upgrader.Upgrade(rw, r, nil)
	utill.HandleErr(err)
}
