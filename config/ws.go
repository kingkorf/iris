package config

import (
	"time"

	"github.com/imdario/mergo"
)

// Currently only these 5 values are used for real
const (
	// DefaultWriteTimeout 10 * time.Second
	DefaultWriteTimeout = 10 * time.Second
	// DefaultPongTimeout 60 * time.Second
	DefaultPongTimeout = 60 * time.Second
	// DefaultPingPeriod (DefaultPongTimeout * 9) / 10
	DefaultPingPeriod = (DefaultPongTimeout * 9) / 10
	// DefaultMaxMessageSize 1024
	DefaultMaxMessageSize = 1024
	// DefaultAutoStart true
	DefaultAutoStart = true
)

//

// Ws the config contains options for 'ws' package
type Ws struct {
	// WriteTimeout time allowed to write a message to the connection.
	WriteTimeout time.Duration
	// PongTimeout allowed to read the next pong message from the connection
	PongTimeout time.Duration
	// PingPeriod send ping messages to the connection with this period. Must be less than PongTimeout
	PingPeriod time.Duration
	// MaxMessageSize max message size allowed from connection
	MaxMessageSize int
	// AutoStart starts the websocket server automatically
	AutoStart bool
}

// DefaultWs returns the default config for iris-ws
func DefaultWs() Ws {
	return Ws{
		WriteTimeout:   DefaultWriteTimeout,
		PongTimeout:    DefaultPongTimeout,
		PingPeriod:     DefaultPingPeriod,
		MaxMessageSize: DefaultMaxMessageSize,
		AutoStart:      DefaultAutoStart,
	}
}

// Merge merges the default with the given config and returns the result
func (c Ws) Merge(cfg []Ws) (config Ws) {

	if len(cfg) > 0 {
		config = cfg[0]
		mergo.Merge(&config, c)
	} else {
		_default := c
		config = _default
	}

	return
}

// MergeSingle merges the default with the given config and returns the result
func (c Ws) MergeSingle(cfg Ws) (config Ws) {

	config = cfg
	mergo.Merge(&config, c)

	return
}
