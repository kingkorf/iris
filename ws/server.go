package ws

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/websocket"
)

type (
	// ConnectionFunc is the callback which fires when a client/connection is connected to the server.
	// Receives one parameter which is the Connection
	ConnectionFunc func(Connection)

	// Server is the websocket server
	Server interface {
		// Upgrade upgrades the client in order websocket works
		Upgrade(context.IContext) error
		// OnConnection registers a callback which fires when a connection/client is connected to the server
		OnConnection(ConnectionFunc)
	}

	server struct {
		upgrader              websocket.Upgrader
		put                   chan *connection
		free                  chan *connection
		connections           map[*connection]bool
		messages              chan messagePayload
		onConnectionListeners []ConnectionFunc
		//connectionPool        *sync.Pool // sadly I can't make this because the websocket connection is live until is closed.
	}
)

var _ Server = &server{}

// server implementation

func newServer(autostart bool) *server {
	s := &server{
		put:                   make(chan *connection),
		free:                  make(chan *connection),
		connections:           make(map[*connection]bool),
		messages:              make(chan messagePayload),
		onConnectionListeners: make([]ConnectionFunc, 0),
	}

	s.upgrader = websocket.New(s.handleConnection)
	if autostart {
		go s.serve()
	}
	return s
}

func (s *server) Upgrade(ctx context.IContext) error {
	return s.upgrader.Upgrade(ctx)
}

func (s *server) handleConnection(websocketConn *websocket.Conn) {
	c := newConnection(websocketConn, s.messages, s.free)
	s.put <- c
	go c.writer()
	c.reader()
}

func (s *server) OnConnection(cb ConnectionFunc) {
	s.onConnectionListeners = append(s.onConnectionListeners, cb)
}

func (s *server) serve() {
	for {
		select {
		case c := <-s.put: // connection connected
			s.connections[c] = true
			for i := range s.onConnectionListeners {
				s.onConnectionListeners[i](c)
			}
		case c := <-s.free: // connection closed
			if _, found := s.connections[c]; found {
				delete(s.connections, c)
			}
		case msg := <-s.messages: // message received from the connection

			for c := range s.connections {
				if msg.to != All { // if it's not suppose to send to all connections (including itself)
					if msg.to != NotMe && msg.to != c.id { // if to specific To
						continue
					} else if msg.to == NotMe && msg.from == c.id { // if broadcast to other connections except this
						continue
					}
				}
				select {
				case c.send <- msg.data: //send the message back to the connection in order to send it to the client
				default:
					c.Disconnect()
				}

			}
		}

	}
}

//
