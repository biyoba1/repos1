package handler

import (
	"github.com/biyoba1/statistic_service/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	orderBook := router.Group("/orderBook")
	{
		orderBook.GET("/", h.getOrderBook)
		orderBook.POST("/", h.saveOrderBook)
	}

	orderHistory := router.Group("orders")
	{
		orderHistory.GET("/", h.getOrders)
		orderHistory.POST("/", h.saveOrder)
	}

	return router
}
