# Package information

This package is new and unique, if you notice a bug or issue [post it here](https://github.com/kataras/iris/issues).

It is not fully ready yet, it doesn't contains all the features I wanted to implement, yet.

# Usage

But it's working for basics, runnable example can be found [here](https://github.com/iris-contrib/examples/tree/master/websocket).



**Server-side**
```go
import "github.com/kataras/iris/websocket"
//...

// important staff

w := websocket.New(api, "/my_endpoint")
// for default 'iris.' station use that: w := websocket.New(iris.DefaultIris, "/my_endpoint")

w.OnConnection(func(c websocket.Connection) {
	c.On("chat", func(message string) {
		c.To(websocket.Broadcast).Emit("chat", "Message from: "+c.ID()+"-> "+message) // to all except this connection
		// c.To("to a specific connection.ID() [rooms are coming soon]").Emit...
		c.Emit("chat", "Message from myself: "+message)
	})
})

//


// ...

```


**Client-side**

```js
// js/chat.js
var messageTxt;
var messages;

$(function () {

	messageTxt = $("#messageTxt");
	messages = $("#messages");


	ws = new Ws("ws://" + HOST + "/my_endpoint");
	ws.OnConnect(function () {
		console.log("Websocket connection enstablished");
	});

	ws.OnDisconnect(function () {
		appendMessage($("<div><center><h3>Disconnected</h3></center></div>"));
	});

	ws.On("chat", function (message) {
		appendMessage($("<div>" + message + "</div>"));
	})

	$("#sendBtn").click(function () {
		//ws.EmitMessage(messageTxt.val());
		ws.Emit("chat", messageTxt.val().toString());
		messageTxt.val("");
	})

})


function appendMessage(messageDiv) {
    var theDiv = messages[0]
    var doScroll = theDiv.scrollTop == theDiv.scrollHeight - theDiv.clientHeight;
    messageDiv.appendTo(messages)
    if (doScroll) {
        theDiv.scrollTop = theDiv.scrollHeight - theDiv.clientHeight;
    }
}

```


```html

<html>

<head>
	<title>My iris-ws</title>
</head>

<body>
	<div id="messages" style="border-width:1px;border-style:solid;height:400px;width:375px;">

	</div>
	<input type="text" id="messageTxt" />
	<button type="button" id="sendBtn">Send</button>
	<script type="text/javascript">
		var HOST = {{.Host}}
	</script>
	<script src="js/vendor/jquery-2.2.3.min.js" type="text/javascript"></script>
	<!-- /iris-ws.js is served automatically by the server -->
	<script src="/iris-ws.js" type="text/javascript"></script>
	<!-- -->
	<script src="js/chat.js" type="text/javascript"></script>
</body>

</html>


```

**Server-Side second approach with rooms**


Let's use the rooms feature to achieve the same results, no need to modify client, just the **server-side** :

```go
package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)

type clientPage struct {
	Title string
	Host  string
}

func main() {
	api := iris.New()

	api.Static("/js", "./static/js", 1)

	api.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.HostString()})
	})

	// important staff

	w := websocket.New(api, "/my_endpoint")
	// for default 'iris.' station use that: w := websocket.New(iris.DefaultIris, "/my_endpoint")

	var myChatRoom = "room1" // here the room name, you can define anything you want, you can make an event to
	//  the client to specify which room want join and so on

	w.OnConnection(func(c websocket.Connection) {

		w.Join(myChatRoom, c) // join this connection to the "room1"

		c.On("chat", func(message st****ring) {
			//send the message to the whole room, all connections are inside this room will receive this message
			c.To(myChatRoom).Emit("chat", "From: "+c.ID()+": "+message)
		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})
	})

	//

	fmt.Println("Server is listening at: 8080")
	api.Listen(":8080")
}


```


[E-Book section](https://kataras.gitbooks.io/iris/content/package-websocket.html)
