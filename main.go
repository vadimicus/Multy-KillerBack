package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

const (
	PORT = ":5555"
	room_wireless = "wireless"
	event_connection = "connection"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On(event_connection, func(so socketio.Socket) {
		log.Println("on connection")
		log.Printf("So id:", so.Id())
		log.Printf("So request", so.Request().GetBody)
		so.Join(room_wireless)
		//log.Printf("Some data:", so.Id(), so.Request(), so)


		so.On("chat message", func(msg string) {
			log.Printf("Some data:", so.Id(), so.Request(), so)
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			log.Printf("So id:", so.Id())
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Printf("Some data:", so.Id(), so.Request(), so)
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	//http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost" + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}