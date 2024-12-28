package controllers

import (
	"golangcrud/database"
	"golangcrud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePayment
func CreatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order models.Order
	if err := database.DB.First(&order, payment.OrderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Mark order
	payment.Status = "paid"
	order.Status = "completed"

	database.DB.Save(&order)
	database.DB.Create(&payment)
	c.JSON(http.StatusCreated, payment)
}
