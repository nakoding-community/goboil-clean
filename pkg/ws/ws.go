package ws

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func NewWs(c echo.Context) error {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	ch := NewChannel(ws)
	go wsReceiver(ch)
	go wsSender(ch)

	<-ch.Done
	logrus.Println("NewWs: done")
	return nil
}

func wsReceiver(ch *Channel) {
	defer func() {
		ch.Ws.Close()
	}()

	for {
		msgType, msg, err := ch.Ws.ReadMessage()
		if err != nil {
			break
		}
		logrus.Println("wsReceiver:", msgType, string(msg))
		ch.Command <- Command{
			MsgType: msgType,
			Msg:     msg,
		}
	}
	close(ch.Done)
}

func wsSender(ch *Channel) {
	defer func() {
		ch.Ws.Close()
	}()

breakLoop:
	for {
		select {
		case v := <-ch.Command:
			err := ch.Ws.WriteMessage(v.MsgType, []byte(v.Msg))
			if err != nil {
				logrus.Println(err)
				break breakLoop
			}
		case <-ch.Done:
			return
		}
	}
	close(ch.Done)
}
