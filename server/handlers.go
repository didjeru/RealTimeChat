package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// ApplyHandlers apply hadlers for server
func (serv *Server) ApplyHandlers() {
	serv.router.Handle("/*", http.FileServer(http.Dir("./web")))
	serv.router.Get("/socket", serv.socketHandler)
}

func (serv *Server) socketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := serv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("websocket err: %v", err)
	}
	stopchan := make(chan string)

	go func() {
		for {
			<-time.After(5 * time.Second)
			if <-stopchan == "exit" {
				break
			}
			msg := Message{
				Type: MTPing,
			}
			if err := ws.WriteJSON(msg); err != nil {
				log.Printf("ws send ping err: %v", err)
			}
		}
	}()

	id := uuid.New().String()
	user := NewUser(id, ws)
	channels, _ := json.Marshal(serv.publisher.GetChannels())
	ws.WriteJSON(Message{
		Type: MTChannels,
		Data: string(channels),
	})

	for {
		msg := Message{}
		if err := ws.ReadJSON(&msg); err != nil {
			if !websocket.IsCloseError(err, 1001) {
				log.Fatalf("ws msg read err: %v", err)
			}
			break
		}

		if msg.Type == MTJoin {
			serv.channels[msg.Channel].Subscribe(user)
		}

		if msg.Type == MTPong {
			continue
		}

		if msg.Type == MTMessage {
			fmt.Println(msg.Data, msg.Channel)
			serv.publisher.Send(msg.Channel+": "+"("+user.GetID()+") "+msg.Data, msg.Channel)
		}

		if msg.Type == MTLeave {
			fmt.Println(msg.Channel)
			serv.channels[msg.Channel].UnSubscribe(user)
		}

	}

	fmt.Println("User leaving...")
	defer func() {
		for _, v := range serv.channels {
			v.UnSubscribe(user)
		}
		stopchan <- "exit"
	}()
}
