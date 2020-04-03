package server

import (
	"goPat/realtimechat/event_channel"
	"log"

	"github.com/gorilla/websocket"
)

// User struct for websocket user
type User struct {
	event_channel.SubscriberDefault
	Username string
	ws       *websocket.Conn
}

// NewUser create new websocket user
func NewUser(username string, ws *websocket.Conn) *User {
	return &User{
		Username: username,
		ws:       ws,
	}
}

// OnReceive receive new message
func (u *User) OnReceive(msg string) {
	m := Message{
		Type: MTMessage,
		Data: msg,
	}
	if err := u.ws.WriteJSON(m); err != nil {
		log.Printf("ws msg fetch err: %v", err)
	}
}

// GetID return user ID
func (u *User) GetID() string {
	return u.Username
}
