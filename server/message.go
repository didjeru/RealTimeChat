package server

//MessageType is message type string
type MessageType string

const (
	//MTPing is meesage ping
	MTPing MessageType = "ping"
	//MTPong is message pong
	MTPong MessageType = "pong"
	//MTMessage is message
	MTMessage MessageType = "message"
)

//Message is message struct
type Message struct {
	Type MessageType `json:"type"`
	Data string      `json:"data,omitempty"`
}
