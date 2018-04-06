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
	EVENT_NEW_RECEIVER	 = "event:new:receiver"
	//EVENT_RECEIVER_ON_OK = "event:receiver:on:ok"
	EVENT_RECEIVER_OFF = "event:receiver:off"
	EVENT_SENDER_ON = "event:sender:on"
	EVENT_SENDER_UPDATE = "event:sender:update"
)




type Receiver struct {
	Id			string		`json:"user_id"`
	CurrencyId	int			`json:"currency_id"`
	Amount		int64		`json:"amount"`
	UserCode	int			`json:"user_code"`

}

type Sender struct {
	Id			string		`json:"user_id"`
	UserCode	int			`json:"user_code"`

}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	receivers := make(map[int]Receiver)
	senders := []Sender{}


	server.On(EVENT_CONNECTION, func(so socketio.Socket) {
		log.Println("on connection")
		log.Printf("So id:", so.Id())
		log.Printf("So request", so.Request().GetBody)

		so.Join(ROOM_WIRELESS)
		so.BroadcastTo(ROOM_WIRELESS, EVENT_NEW_RECEIVER, "HEYYYS can brodcast")
		//log.Printf("Some data:", so.Id(), so.Request(), so)

		so.On(EVENT_RECEIVER_ON, func(data interface{} ) string {
			log.Printf("So id Event RECEIVER ON:", so.Id())
			log.Printf("DAta:", data)
			//so.Emit(EVENT_RECEIVER_ON_OK, "nice job maaan ;)")
			so.Join(so.Id())

			parsed := data.(map[string]interface{})
			user_id:= parsed["user_id"].(string)
			raw_code:= parsed["user_code"].(float64)
			user_code := int(raw_code)
			raw_currency_id := parsed["currency_id"].(float64)
			currency_id := int(raw_currency_id)
			raw_amount := parsed["amount"].(float64)
			amount := int64(raw_amount)

			receiver := Receiver{
				Id:user_id,
				UserCode:user_code,
				CurrencyId:currency_id,
				Amount:amount,
			}


			_, ok := receivers[receiver.UserCode]
			if !ok{
				receivers[receiver.UserCode]=receiver
			}

			//Try to find Sender
			for _,sender := range senders{
				if sender.UserCode == receiver.UserCode{
					so.BroadcastTo(ROOM_WIRELESS, EVENT_NEW_RECEIVER, "HEY")
				}
			}

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

			//map[string]interface {}=map[user_id:Vadddim user_code:3252])
			parsed := data.(map[string]interface{})
			user_id:= parsed["user_id"].(string)
			raw_code:= parsed["user_code"].(float64)
			user_code := int(raw_code)




			log.Printf("God data from sender:", user_id, user_code)



			sender:= Sender{UserCode:user_code, Id:user_id }

			so.BroadcastTo(ROOM_WIRELESS, EVENT_NEW_RECEIVER, "HEyyyy")

			var senderExist bool = false

			for _, cachedSender:= range senders{
				if cachedSender.Id == sender.Id{
					senderExist = true
				}
			}


			if !senderExist{
				senders = append(senders, sender)
			}

			// try to find Receiver by the code
			receiver, ok := receivers[sender.UserCode]
			if ok{
				receiver.Id = "new ID"
				so.BroadcastTo(ROOM_WIRELESS, EVENT_NEW_RECEIVER, "Heeey")
			} else{
				var senderExist bool = false

				for _, cachedSender:= range senders{
					if cachedSender.Id == sender.Id{
						senderExist = true
					}
				}


				if !senderExist{
					senders = append(senders, sender)
				}


			}

			return "hello sender"
		})


		//so.On(EVENT_SENDER_OFF, func(data interface{} ) string {
		//	log.Printf("Sender become off:", so.Id())
		//	//TODO add sender disconnect logic here
		//	return "goodbye sender"
		//})
		//
		//so.On(EVENT_SENDER_UPDATE, func(data interface{} ) string {
		//	log.Printf("Sender update called:", so.Id(), data)
		//	//TODO add sender update logic here
		//	return "goodbye sender"
		//})


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
