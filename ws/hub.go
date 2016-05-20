package ws

/* Notes for me
1. I though to make it like .On("connection", func(c *Connection){}) but I decided to make it more statically typed .OnConnect(func(c *Connection){}).
2. The internal events will fire on with the .fire$Event$.
3. The external-websocket events will fire with .Emit$Event manually by the developer


*/

const (
	PayloadText        = 0
	PayloadBinary      = 1
	DefaultPayloadType = PayloadText
)

type (
	PayloadType uint8

	Payload struct {
		From string // from the connetion id
		To   string // empty means all
		Type PayloadType
	}

	Message struct {
		payload Payload
		data    []byte
	}

	Hub struct {
		// Register adds a connection from the list
		Register chan *Connection
		// Unregister removes a connection from the list
		Unregister chan *Connection
		// Connections the list of registed connectons
		Connections map[*Connection]bool
		// Broadcast messages from the connections
		Broadcast chan Message
		/* events */
		connectListeners []ConnectEventFunc
	}
)

func NewHub() *Hub {
	return &Hub{
		Register:         make(chan *Connection),
		Unregister:       make(chan *Connection),
		Connections:      make(map[*Connection]bool),
		Broadcast:        make(chan Message),
		connectListeners: make([]ConnectEventFunc, 0),
	}
}

/***** *****/
func (h *Hub) OnConnection(cb ConnectEventFunc) {
	h.connectListeners = append(h.connectListeners, cb)
}

func (h *Hub) fireConnect(c *Connection) {
	for i := range h.connectListeners {
		h.connectListeners[i](c)
	}
}

/***** *****/

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.Register:
			h.Connections[c] = true
			h.fireConnect(c)
		case c := <-h.Unregister:
			if _, ok := h.Connections[c]; ok {
				delete(h.Connections, c)
				close(c.send)
				c.fireDisconnect()
			}

		case message := <-h.Broadcast:
			for c := range h.Connections {
				to := message.payload.To
				if to != ToAll { // if it's not suppose to send to all connections (including itself)
					if to != ToAllExceptMe && to != c.Id { // if to specific To, connection.To("connectionid").Emit([]byte("msg"))
						println("Hub:89 break1")
						continue

					} else if to == ToAllExceptMe && message.payload.From == c.Id { // if connection.Broadcast.Emit([]byte("msg"))
						println("Hub:93 break2")
						continue
					}
				}
				select {
				case c.send <- message.data:
				default:
					close(c.send)
					delete(h.Connections, c)
					c.fireDisconnect()
				}
			}
		}
	}
}
