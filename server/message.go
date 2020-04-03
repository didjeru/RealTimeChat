package server

// MessageType is message type string
type MessageType string

const (
	// MTPing message ping
	MTPing MessageType = "ping"
	// MTPong message pong
	MTPong MessageType = "pong"
	// MTMessage message
	MTMessage MessageType = "message"
	// MTJoin join message
	MTJoin MessageType = "join"
	// MTChannels list of channels
	MTChannels MessageType = "channels"
	//MTLeave leave channel
	MTLeave MessageType = "leave"
)

// Message is message struct
type Message struct {
	Type    MessageType `json:"type"`
	Data    string      `json:"data,omitempty"`
	Channel string      `json:"channel,omitempty"`
}
