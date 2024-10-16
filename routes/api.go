package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rudSarkar/pubsub-websocket/config"
	"github.com/rudSarkar/pubsub-websocket/helper"
	"github.com/rudSarkar/pubsub-websocket/migrations"
	"github.com/rudSarkar/pubsub-websocket/model"
)

// Create new order and push to broadcast
func CreateOrder(c *gin.Context) {
	var order migrations.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Insert the new order into the database
	if result := config.DB.Create(&order); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}
	// Log the order being sent to the WebSocket channel
	log.Println("Pushing latest order to broadcast:", order)

	// Push the new order to the 'orders' topic
	helper.Broadcast <- helper.Message{
		Topic: "orders",
		Data:  order,
	}
	c.JSON(http.StatusOK, order)
}

func GetLatestOrder(c *gin.Context) {
	var orders []model.Order
	if result := config.DB.Order("id desc").Find(&orders); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No orders found"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// Create new bill and push to broadcast
func CreateBill(c *gin.Context) {
	var bill migrations.Bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Insert the new bill into the database
	if result := config.DB.Create(&bill); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bill"})
		return
	}

	// Log the bill being sent to the WebSocket channel
	log.Println("Pushing latest bill to broadcast:", bill)

	// Push the new bill to the 'bills' topic
	helper.Broadcast <- helper.Message{
		Topic: "bills",
		Data:  bill,
	}
	c.JSON(http.StatusOK, bill)
}

func GetLatestBill(c *gin.Context) {
	var bills []model.Bill
	if result := config.DB.Order("id desc").Find(&bills); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No bills found"})
		return
	}

	c.JSON(http.StatusOK, bills)
}
