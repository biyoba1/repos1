package handler

import (
	statTracker "github.com/biyoba1/statistic_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getOrders(c *gin.Context) {
	clientName := c.Query("client_name")
	client := &statTracker.Client{
		Client_name: clientName,
	}
	orders, err := h.services.GetOrderHistory(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *Handler) saveOrder(c *gin.Context) {
	var req struct {
		Client statTracker.Client
		Order  statTracker.HistoryOrder
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.SaveOrder(&req.Client, &req.Order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order saved successfully"})
}
