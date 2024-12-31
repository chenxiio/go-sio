package main

import (
	"fmt"
	"log"

	gosio "github.com/gnabgib/go-sio"
	"github.com/gnabgib/go-sio/transport"
)

func main() {
	//connect to server, you can use your own transport settings
	parms := make(map[string]string)
	tr := transport.GetDefaultWebsocketTransport()
	ws := gosio.New(gosio.GetURL("localhost", 10600, false, &parms), tr)

	ws.OnDisconnect(func(c *gosio.Channel) {
		log.Println("Disconnected to server1")

		err := ws.Dial()
		if err != nil {

			fmt.Printf("dial err %s", err.Error())
		}

	})

	ws.OnConnect(func(c *gosio.Channel) {
		log.Println("Connected to server1")
		ws.Emit("subio", "IO")
	})

	ws.On("subioIO", func(c *gosio.Channel, msg map[string]any) {

		log.Println("Received Devicevalues:", c.ID(), msg)
	})

	ws.Dial()
	//Do something with the websocket
	//	ws.Emit("chat message", "hi")
	select {}
}
