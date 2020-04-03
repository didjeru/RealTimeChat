package server

import (
	"goPat/realtimechat/event_channel"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

//Server struct
type Server struct {
	router   *chi.Mux
	upgrader *websocket.Upgrader

	channels  map[string]*event_channel.Channel
	publisher *event_channel.Publisher
}

//New start new server
func New() *Server {
	router := chi.NewRouter()

	ch1 := event_channel.NewChannel()
	ch2 := event_channel.NewChannel()
	ch3 := event_channel.NewChannel()
	ch4 := event_channel.NewChannel()

	pub := event_channel.NewPublisher()

	pub.AddChannel("#ch1", ch1)
	pub.AddChannel("#ch2", ch2)
	pub.AddChannel("#ch3", ch3)
	pub.AddChannel("#ch4", ch4)

	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	serv := &Server{
		router:    router,
		upgrader:  upgrader,
		channels:  map[string]*event_channel.Channel{"#ch1": ch1, "#ch2": ch2, "#ch3": ch3, "#ch4": ch4},
		publisher: pub,
	}

	serv.ApplyHandlers()

	return serv
}

//Start server and return error
func (serv *Server) Start() error {
	return http.ListenAndServe(":8080", serv.router)
}
