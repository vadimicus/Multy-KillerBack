package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

const (
	PORT = ":5555"
	ROOM_WIRELESS = "wireless"
	EVENT_CONNECTION = "connection"
	EVENT_RECEIVER_ON = "event:receiver:on"
	EVENT_RECEIVER_ON_OK = "event:receiver:on:ok"
	EVENT_RECEIVER_OFF = "event:receiver:off"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On(EVENT_CONNECTION, func(so socketio.Socket) {
		log.Println("on connection")
		log.Printf("So id:", so.Id())
		log.Printf("So request", so.Request().GetBody)
		so.Join(ROOM_WIRELESS)
		//log.Printf("Some data:", so.Id(), so.Request(), so)


		so.On(EVENT_RECEIVER_ON, func(data interface{} ){
			log.Printf("So id Event RECEIVER ON:", so.Id())
			log.Printf("DAta:", data)
			so.Emit(EVENT_RECEIVER_ON_OK, "nice job maaan ;)")

		})

		//so.On("chat message", func(msg string) {
		//	log.Printf("Some data:", so.Id(), so.Request(), so)
		//	log.Println("emit:", so.Emit("chat message", msg))
		//	so.BroadcastTo("chat", "chat message", msg)
		//})
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