package websocket

import (
	"kimparty/internal/party"
	"kimparty/internal/party/message"
	"kimparty/internal/party/user"
	"log"
	"time"

	"github.com/lxzan/gws"
)

const (
	PingInterval         = 5 * time.Second
	HeartbeatWaitTimeout = 10 * time.Second
)

type WebSocket struct{}

func (ws *WebSocket) getParty(socket *gws.Conn) *party.Party {
	pt, _ := socket.Session().Load("party")
	return pt.(*party.Party)
}

func (ws *WebSocket) getUser(socket *gws.Conn) *user.User {
	u, _ := socket.Session().Load("user")
	return u.(*user.User)
}

func (ws *WebSocket) OnOpen(socket *gws.Conn) {
	u := ws.getUser(socket)
	pt := ws.getParty(socket)

	userJson, _ := u.ToJSON()
	socket.WriteMessage(gws.OpcodeText, userJson)

	log.Printf("User [%s - %s]: connected to Party [%s]", u.Name, u.ID, pt.ID)
}

func (ws *WebSocket) OnClose(socket *gws.Conn, err error) {
	pt := ws.getParty(socket)
	u := ws.getUser(socket)

	pt.RemoveConn(u.ID)

	leavedMsg, _ := message.NewLeavedMessage(pt.ID, u.ID)
	broadcastMessage(pt, leavedMsg)

	log.Printf("User [%s - %s]: disconnected from Party [%s]", u.Name, u.ID, pt.ID)
}

func (ws *WebSocket) OnError(socket *gws.Conn, err error) {
	pt := ws.getParty(socket)
	u := ws.getUser(socket)

	pt.RemoveConn(u.ID)

	log.Printf("User [%s - %s]: disconnected from Party [%s]: %s", u.Name, u.ID, pt.ID, err)
}

func (ws *WebSocket) OnMessage(socket *gws.Conn, msg *gws.Message) {
	defer msg.Close()
	bytes := msg.Data.Bytes()

	if len(bytes) == 4 && string(bytes) == "ping" {
		ws.OnPing(socket, nil)
		return
	}

	pt := ws.getParty(socket)
	userID := ws.getUser(socket).ID

	msgJson, err := message.NewFromUserMessage(pt.ID, userID, bytes)

	if err != nil {
		socket.WriteString(err.Error())
		return
	}

	broadcastMessage(pt, msgJson)
}

func (c *WebSocket) OnPing(socket *gws.Conn, payload []byte) {
	socket.SetDeadline(time.Now().Add(PingInterval + HeartbeatWaitTimeout))
	socket.WriteString("pong")
}

func (c *WebSocket) OnPong(socket *gws.Conn, payload []byte) {}

func broadcastMessage(pt *party.Party, msg []byte) {
	bc := gws.NewBroadcaster(gws.OpcodeText, msg)
	defer bc.Close()

	for socket := range pt.IterConns() {
		bc.Broadcast(socket)
	}
}
