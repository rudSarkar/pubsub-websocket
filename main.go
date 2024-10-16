package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudSarkar/pubsub-websocket/config"
	"github.com/rudSarkar/pubsub-websocket/helper"
	"github.com/rudSarkar/pubsub-websocket/middleware"
	"github.com/rudSarkar/pubsub-websocket/routes"
)

func main() {
	r := gin.Default()

	// Add CORS
	r.Use(middleware.CORSMiddleware())

	// Initialize database and auto-migrate tables
	config.InitDB()
	config.MigrateDB()

	// WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		helper.HandleConnections(c)
	})

	// Order and Bill endpoints for creating new order and bill
	r.POST("/order", routes.CreateOrder)
	r.POST("/bill", routes.CreateBill)

	// Order and Bill endpoint for fetching latest order and bill
	r.GET("/order", routes.GetLatestOrder)
	r.GET("/bill", routes.GetLatestBill)

	// Start broadcasting messages
	go helper.HandleMessages()

	r.Run(":8080")
}
