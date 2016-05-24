package ws

const (
	// All is the string which the Emmiter use to send a message to all
	All = ""
	// NotMe is the string which the Emmiter use to send a message to all except this connection
	NotMe = ";iris;to;all;except;me;"
	// Broadcast is the string which the Emmiter use to send a message to all except this connection, same as 'NotMe'
	Broadcast = NotMe
)

type (
	// Emmiter is the message/or/event manager
	Emmiter interface {
		// EmitMessage sends a native websocket message
		EmitMessage(string) error
		// Emit sends a message on a particular event
		Emit(string, interface{}) error
	}

	messagePayload struct {
		from string
		to   string
		data []byte
	}

	emmiter struct {
		conn *connection
		to   string
	}
)

var _ Emmiter = &emmiter{}

// payload implementation

func newMessagePayload(from string, to string, data []byte) messagePayload {
	return messagePayload{from, to, data}
}

//

// emmiter implementation

func newEmmiter(c *connection, to string) *emmiter {
	return &emmiter{conn: c, to: to}
}

func (e *emmiter) EmitMessage(nativeMessage string) error {
	mp := newMessagePayload(e.conn.id, e.to, []byte(nativeMessage))
	e.conn.messages <- mp
	return nil
}

func (e *emmiter) Emit(event string, data interface{}) error {
	message, err := encodeMessage(event, data)
	if err != nil {
		return err
	}
	e.EmitMessage(message)
	return nil
}

//
