package ws

const (
	// for future use when dynamic alternative use .On("anything",interface{})
	ConnectEvent    = "connection"
	DisconnectEvent = "disconnect"
	MessageEvent    = "message"
)

type (
	// Hub
	ConnectEventFunc func(*Connection)
	// Connection
	DisconnectEventFunc func()
	MessageEventFunc    func(message []byte)
)
