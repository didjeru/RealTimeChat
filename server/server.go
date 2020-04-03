package server

import (
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

//Subscriber func with message
type Subscriber func(msg string) error

//Server struct
type Server struct {
	router   *chi.Mux
	upgrader *websocket.Upgrader

	submutex    *sync.Mutex
	subscribers map[string]Subscriber
}

//New start new server
func New() *Server {
	router := chi.NewRouter()

	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	serv := &Server{
		router:   router,
		upgrader: upgrader,

		submutex:    &sync.Mutex{},
		subscribers: map[string]Subscriber{},
	}

	serv.ApplyHandlers()

	return serv
}

//Start server and return error
func (serv *Server) Start() error {
	return http.ListenAndServe(":8080", serv.router)
}
