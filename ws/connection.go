package ws

import (
	"time"

	"github.com/kataras/iris/utils"
	"github.com/kataras/iris/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512

	ToAll         = ""
	ToAllExceptMe = "_TO_ALL_NO_ME"
)

type (
	iemmiter interface {
		Emit([]byte)
	}

	emmiter struct {
		conn *Connection
		to   string
	}

	Connection struct {
		Id     string
		wsConn *websocket.Conn
		/* internal events to -> ws-> hub */
		NotifyClose   func()
		NotifyMessage func(Message)
		/**/
		send chan []byte

		/* events */
		disconnectListeners []DisconnectEventFunc
		messageListeners    []MessageEventFunc

		/* emitters */
		Broadcast iemmiter
	}
)

func (e emmiter) Emit(data []byte) {
	message := Message{payload: Payload{From: e.conn.Id, To: e.to, Type: DefaultPayloadType}, data: data}
	e.conn.NotifyMessage(message) // -> hub
}

// NewConnection creates a connection and returns it
func NewConnection(websocketConnection *websocket.Conn) *Connection {
	c := &Connection{
		Id:                  utils.RandomString(64),
		wsConn:              websocketConnection,
		send:                make(chan []byte, 256),
		disconnectListeners: make([]DisconnectEventFunc, 0),
		messageListeners:    make([]MessageEventFunc, 0),
	}
	c.Broadcast = emmiter{conn: c, to: ToAllExceptMe}
	return c
}

func (c *Connection) To(to string) iemmiter {
	return emmiter{conn: c, to: to}
}

func (c *Connection) Emit(data []byte) {
	//same as To(c.Id).Emit
	c.To(c.Id).Emit(data)
}

/***** ****/

func (c *Connection) OnDisconnect(cb DisconnectEventFunc) {
	c.disconnectListeners = append(c.disconnectListeners, cb)
}

func (c *Connection) OnMessage(cb MessageEventFunc) {
	c.messageListeners = append(c.messageListeners, cb)
}

func (c *Connection) fireDisconnect() {
	for i := range c.disconnectListeners {
		c.disconnectListeners[i]()
	}
}

func (c *Connection) fireMessage(message []byte) {
	for i := range c.messageListeners {
		c.messageListeners[i](message)
	}
}

/***** ****/

/*** ***/

/*** ***/
func (c *Connection) Listen() {
	go c.startWriter()
	c.startReader()
}

func (c *Connection) write(messageType int, payload []byte) error {
	c.wsConn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.wsConn.WriteMessage(messageType, payload)
}

// startWriter sends messages from the hub to the websocket client/connection
func (c *Connection) startWriter() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.NotifyClose()
		c.wsConn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *Connection) startReader() {
	defer func() {
		// if this breaks means that an error happen
		c.NotifyClose() // send notification to the hub and all listeners that this conn is closed by an error or timeout
		c.wsConn.Close()
	}()

	c.wsConn.SetReadLimit(maxMessageSize)
	c.wsConn.SetReadDeadline(time.Now().Add(pongWait))
	c.wsConn.SetPongHandler(func(string) error {
		c.wsConn.SetReadDeadline(time.Now().Add(pongWait)) // on pong just continue the connection by extend its life
		return nil
	})

	for {
		_, data, err := c.wsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				println(err.Error())
			}
			break
		}

		c.fireMessage(data)
	}
}
