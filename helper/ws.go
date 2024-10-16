package helper

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client structure that includes the WebSocket connection and its subscribed topic
type Client struct {
	Conn  *websocket.Conn
	Topic string
}

// Clients map will store the clients and their subscribed topics
var Clients = make(map[*Client]bool)

// Channel to broadcast messages to topics
var Broadcast = make(chan Message)

// Message structure to hold the topic and data
type Message struct {
	Topic string
	Data  interface{}
}

// Handle WebSocket connection
func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	client := &Client{Conn: ws}

	// Read the subscription message to know the topic
	var subscribeMsg struct {
		Topic string `json:"topic"`
	}

	err = ws.ReadJSON(&subscribeMsg)
	if err != nil {
		log.Printf("Subscription error: %v", err)
		return
	}

	// Assign the topic to the client
	client.Topic = subscribeMsg.Topic
	Clients[client] = true

	log.Printf("New client subscribed to topic: %s", client.Topic)

	// Listen for messages from the client
	for {
		var msg interface{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(Clients, client)
			break
		}
	}
}

// Broadcast messages to all clients subscribed to a specific topic
func HandleMessages() {
	for {
		msg := <-Broadcast

		log.Printf("Broadcasting message to topic: %s", msg.Topic)

		for client := range Clients {
			// Only send the message to clients subscribed to the right topic
			if client.Topic == msg.Topic {
				err := client.Conn.WriteJSON(msg.Data)
				if err != nil {
					log.Printf("error: %v", err)
					client.Conn.Close()
					delete(Clients, client)
				}
			}
		}
	}
}
