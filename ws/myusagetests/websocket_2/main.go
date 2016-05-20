package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/ws"
)

type clientPage struct {
	Title string
	Host  string
}

func main() {
	w := ws.New()
	chat(w)
	iris.Static("/js", "./static/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.HostString()})
	})

	iris.Get("/ws", func(ctx *iris.Context) {
		if err := w.Do(ctx); err != nil {
			iris.Logger().Panic(err)
		}
	})

	fmt.Println("Server is listening at: 8080")
	iris.Listen(":8080")
}

func chat(w *ws.Ws) {
	w.OnConnection(func(c *ws.Connection) {
		println("connection with Id: " + c.Id + " just connected")
		c.OnMessage(func(message []byte) {
			println("Message sent: " + string(message) + " from: " + c.Id)
			c.Broadcast.Emit([]byte("From: " + c.Id + "-> " + string(message))) // to all except this connection //worked
			//c.To(string(message)).Emit(message) //send the socket id which is the given message //worked
			c.Emit([]byte("to my self: " + string(message))) // worked
		})
	})

}
