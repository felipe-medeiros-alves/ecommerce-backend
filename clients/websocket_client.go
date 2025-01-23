package clients

import (
	"ecommerce-backend/models"
	"ecommerce-backend/services"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		err = services.SendMessage(msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		ws.WriteJSON(msg)
	}
}
