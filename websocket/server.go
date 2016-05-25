package websocket

import (
	"sync"

	"github.com/iris-contrib/websocket"
	"github.com/kataras/iris/context"
)

type (
	// ConnectionFunc is the callback which fires when a client/connection is connected to the server.
	// Receives one parameter which is the Connection
	ConnectionFunc func(Connection)

	Rooms map[string][]string
	// Server is the websocket server
	Server interface {
		// Upgrade upgrades the client in order websocket works
		Upgrade(context.IContext) error
		// OnConnection registers a callback which fires when a connection/client is connected to the server
		OnConnection(ConnectionFunc)

		// Join join a connection to a room, it doesn't check if connection is already there, so care
		Join(string, Connection)
		// Leave leave a connection from a room, returns true if connection was inside this room and left
		Leave(string, Connection) bool
	}

	server struct {
		upgrader              websocket.Upgrader
		put                   chan *connection
		free                  chan *connection
		connections           map[string]*connection
		rooms                 Rooms // by default a connection is joined to a room which has the connection id as its name
		messages              chan messagePayload
		onConnectionListeners []ConnectionFunc
		mu                    sync.Mutex // for rooms
		//connectionPool        *sync.Pool // sadly I can't make this because the websocket connection is live until is closed.
	}
)

var _ Server = &server{}

// server implementation

func newServer(autostart bool) *server {
	s := &server{
		put:                   make(chan *connection),
		free:                  make(chan *connection),
		connections:           make(map[string]*connection),
		rooms:                 make(map[string][]string),
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

func (s *server) join(roomName string, connID string) {
	s.mu.Lock()
	if s.rooms[roomName] == nil {
		s.rooms[roomName] = make([]string, 0)
	}
	s.rooms[roomName] = append(s.rooms[roomName], connID)
	s.mu.Unlock()
}

func (s *server) Join(roomName string, conn Connection) {
	s.join(roomName, conn.ID())

}

func (s *server) leave(roomName string, connID string) bool {
	s.mu.Lock()
	if s.rooms[roomName] != nil {
		for i := range s.rooms[roomName] {
			if s.rooms[roomName][i] == connID {
				s.rooms[roomName][i] = s.rooms[roomName][len(s.rooms[roomName])-1]
				s.rooms[roomName] = s.rooms[roomName][:len(s.rooms[roomName])-1]
				s.mu.Unlock()
				return true
			}
		}
		if len(s.rooms[roomName]) == 0 { // if room is empty then delete it
			delete(s.rooms, roomName)
		}
	}

	s.mu.Unlock()
	return false
}

func (s *server) Leave(roomName string, conn Connection) bool {
	return s.leave(roomName, conn.ID())
}

func (s *server) serve() {
	for {
		select {
		case c := <-s.put: // connection connected
			s.connections[c.id] = c
			// join to its own room ( no need to use the s.join for locking here)
			s.rooms[c.id] = make([]string, 0)
			s.rooms[c.id] = append(s.rooms[c.id], c.id)
			for i := range s.onConnectionListeners {
				s.onConnectionListeners[i](c)
			}
		case c := <-s.free: // connection closed
			if _, found := s.connections[c.id]; found {
				// leave from all rooms
				for roomName := range s.rooms {
					if len(s.rooms[roomName]) > 0 { // I know its weird here, because we do range below but believe me its needed for these cases
						for i := range s.rooms[roomName] {
							if s.rooms[roomName][i] == c.id {
								s.rooms[roomName][i] = s.rooms[roomName][len(s.rooms[roomName])-1]
								s.rooms[roomName] = s.rooms[roomName][:len(s.rooms[roomName])-1]
							}
						}
						if len(s.rooms[roomName]) == 0 { // if room is empty then delete it
							delete(s.rooms, roomName)
						}
					}

				}

				delete(s.connections, c.id)
			}

		case msg := <-s.messages: // message received from the connection
			// check for room if msg.to != all && msg.to != notme
			if msg.to != All && msg.to != NotMe && s.rooms[msg.to] != nil {
				println("server.go: 150-> to room: " + msg.to)
				// it suppose to send the message to a room
				for _, connectionIdInsideRoom := range s.rooms[msg.to] {
					s.connections[connectionIdInsideRoom].send <- msg.data //here we send it without need to continue below
				}

			} else { // it suppose to send the message to all opened connections or to all except the sender
				for connID := range s.connections {
					if msg.to != All { // if it's not suppose to send to all connections (including itself)
						if msg.to == NotMe && msg.from == connID { // if broadcast to other connections except this
							continue //here we do the opossite of previous block, just skip this connection when it's suppose to send the message to all connections except the sender
						}
					}
					select {
					case s.connections[connID].send <- msg.data: //send the message back to the connection in order to send it to the client
					default:
						s.connections[connID].Disconnect()
					}

				}
			}
		}

	}
}

//
