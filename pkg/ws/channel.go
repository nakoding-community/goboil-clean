package ws

import "github.com/gorilla/websocket"

// var pingPeriod = time.Second * 5
// var writeWait = time.Second * 5
var MsgTypeText = 1

type Channel struct {
	Ws      *websocket.Conn
	Command chan Command
	Done    chan struct{}
}

type Command struct {
	MsgType int
	Msg     []byte
}

func NewChannel(ws *websocket.Conn) *Channel {
	ch := &Channel{
		Ws:      ws,
		Command: make(chan Command),
		Done:    make(chan struct{}),
	}
	return ch
}
