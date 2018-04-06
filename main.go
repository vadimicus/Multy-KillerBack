package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

const (
	PORT                 = ":5555"
	ROOM_WIRELESS        = "wireless"
	EVENT_CONNECTION     = "connection"
	EVENT_RECEIVER_ON    = "event:receiver:on"
	//EVENT_RECEIVER_ON_OK = "event:receiver:on:ok"
	EVENT_RECEIVER_OFF = "event:receiver:off"
	EVENT_SENDER_ON = "event:sender:on"
	EVENT_SENDER_OFF = "event:sender:off"
	EVENT_SENDER_UPDATE = "event:sender:update"
)




type Receiver struct {
	Id			string		`json:"user_id"`
	CurrencyId	int			`json:"currency_id"`
	Amount		int64		`json:"amount"`
	UserCode	int			`json:"user_code"`
	Socker		*socketio.Socket
}

type Sender struct {
	Id			string		`json:"user_id"`
	UserCode	int			`json:"user_code"`
	Socker		*socketio.Socket
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	//receivers := make(map[int]Receiver)
	//senders := make(map[int]Sender)

	server.On(EVENT_CONNECTION, func(so socketio.Socket) {
		log.Println("on connection")
		log.Printf("So id:", so.Id())
		log.Printf("So request", so.Request().GetBody)
		so.Join(ROOM_WIRELESS)
		so.Emit("hi", "HI JACK!")
		//log.Printf("Some data:", so.Id(), so.Request(), so)

		so.On(EVENT_RECEIVER_ON, func(data interface{} ) string {
			log.Printf("So id Event RECEIVER ON:", so.Id())
			log.Printf("DAta:", data)
			//so.Emit(EVENT_RECEIVER_ON_OK, "nice job maaan ;)")
			//TODO add sender logic here
			return "welcome:to:receiver:side"
		})

		so.On(EVENT_RECEIVER_OFF, func(data interface{} ) string {
			log.Printf("So id Event RECEIVER OFF:", so.Id())
			log.Printf("DAta:", data)
			//so.Emit(EVENT_RECEIVER_ON_OK, "nice job maaan ;)")
			//TODO add receiver off logic here
			return "goodbye receiver"
		})

		so.On(EVENT_SENDER_ON, func(data interface{} ) string {
			log.Printf("Sender become on:", so.Id())

			parsed, ok := data.(Sender)
			if ok{
				log.Printf("God data from sender:", parsed)
			} else {
				log.Printf("Something went wrong from sender:", data)
			}


			//TODO parse client from WS and add it to the map

			//sender := Sender{}



			//senders[client.UserCode]=client
			//TODO try to find Receiver by the code
			//receiver, ok := receivers[sender.UserCode]
			//if ok{
				//TODO send reciver to sender
				//TODO and remove this boolshit
				//receiver.UserCode = client.UserCode
			//} else{
				//there is still noreceiver

			//}

			return "hello sender"
		})


		so.On(EVENT_SENDER_OFF, func(data interface{} ) string {
			log.Printf("Sender become off:", so.Id())
			//TODO add sender disconnect logic here
			return "goodbye sender"
		})

		so.On(EVENT_SENDER_UPDATE, func(data interface{} ) string {
			log.Printf("Sender update called:", so.Id(), data)
			//TODO add sender update logic here
			return "goodbye sender"
		})


		so.On("disconnection", func() {
			log.Printf("So id:", so.Id())
			log.Println("on disconnect")
			//TODO add disconnect socket user logic 
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
