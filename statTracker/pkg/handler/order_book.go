package handler

import (
	statTracker "github.com/biyoba1/statistic_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getOrderBook(c *gin.Context) {
	exchange := c.Query("exchange")
	pair := c.Query("pair")
	if exchange == "" || pair == "" {
		newErrorResponse(c, http.StatusBadRequest, "exchange and pair are required")
		return
	}
	orders, err := h.services.GetOrderBook(exchange, pair)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *Handler) saveOrderBook(c *gin.Context) {
	var req struct {
		ExchangeName string                    `json:"exchangeName"`
		Pair         string                    `json:"pair"`
		OrderBook    []*statTracker.DepthOrder `json:"orderBook"`
	}

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.SaveOrderBook(req.ExchangeName, req.Pair, req.OrderBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order book saved successfully"})
}
