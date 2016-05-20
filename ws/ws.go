package ws

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/websocket"
)

// domain for iris station only the methods we need
type IStation interface {
	Get(string)
}

//
type (
	Ws struct {
		upgrader    websocket.Upgrader
		requestPath string
		hub         *Hub
	}
)

// New returns the ws container, which is running
func New() *Ws {
	w := &Ws{
		hub: NewHub(),
	}
	w.upgrader = websocket.New(w.HandleConnection)
	go w.hub.Run()
	return w
}

// Do called once for every http request
func (w *Ws) Do(ctx context.IContext) error {
	return w.upgrader.Upgrade(ctx)
}

func (w *Ws) HandleConnection(websocketConn *websocket.Conn) {
	conn := NewConnection(websocketConn)
	conn.NotifyClose = func() {
		w.hub.Unregister <- conn
	}
	conn.NotifyMessage = func(message Message) {
		w.hub.Broadcast <- message
	}
	w.hub.Register <- conn
	conn.Listen()
}

/* events | from hub */
func (w *Ws) OnConnection(cb ConnectEventFunc) {
	w.hub.OnConnection(cb)
}
