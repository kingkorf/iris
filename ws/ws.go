package ws

import (
	"sync"

	"github.com/kataras/iris/config"
	"github.com/kataras/iris/context"
)

// New returns a new running websocket server
func New() Server {
	return newServer(config.DefaultAutoStart)
}

// singleton
var once sync.Once
var defaultServer = newServer(false)

// Upgrade upgrades the client in order websocket works
func Upgrade(ctx context.IContext) error {
	once.Do(func() {
		go defaultServer.serve()
	})

	return defaultServer.Upgrade(ctx)
}

// OnConnection registers a callback which fires when a connection/client is connected to the server
func OnConnection(cb ConnectionFunc) {
	defaultServer.OnConnection(cb)
}
