package socket

import (
	"net"
	"time"
)

const (
	// WAITING waiting for connecting
	WAITING uint = iota

	// CONNECTING connecting the host
	CONNECTING

	// CONNECTED  the socket is connected
	CONNECTED

	// CLOSED the socket is closed
	CLOSED
)

const (

	// MREAD socket read error ( the local peer closed )
	MREAD uint = iota

	// MBREAK the remote peer closed
	MBREAK

	// MCLOSE user send the CLOSE signal
	MCLOSE
)

// Client the socket client
type Client struct {
	Status   uint
	Conn     net.Conn
	Option   Option
	Handlers Handlers
}

// Option the socket option
type Option struct {
	Protocol   string        `json:"protocol,omitempty"`  // TCP/UDP
	Reconnect  int           `json:"reconnect,omitempty"` // max times try to reconnect server when connection break (client mode only)
	Host       string        `json:"host,omitempty"`
	Port       string        `json:"port,omitempty"`
	Timeout    time.Duration `json:"timeout,omitempty"` // timeout (seconds)
	BufferSize int           `json:"buffer,omitempty"`  // bufferSize
	KeepAlive  time.Duration `json:"keep,omitempty"`    // -1 not keep alive, 0 keep alive always, keep alive n seconds.
}

// Handlers the socket hanlders
type Handlers struct {
	Data      DataHandler
	Error     ErrorHandler
	Close     CloseHandler
	Connected ConnectedHandler
}

// DataHandler Handler
type DataHandler func([]byte, int) ([]byte, error)

// ErrorHandler Handler
type ErrorHandler func(error)

// CloseHandler Handler
type CloseHandler func([]byte, error) []byte

// ConnectedHandler Handler
type ConnectedHandler func(option Option) error
