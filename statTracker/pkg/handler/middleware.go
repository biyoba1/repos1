package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	exchangeCtx = "exchange"
	pairCtx     = "pair"
)

func getExchange(c *gin.Context) (string, error) {
	exchange, ok := c.Get(exchangeCtx)
	if !ok {
		return "", errors.New("exchange not found")
	}

	exchangeString, ok := exchange.(string)
	if !ok {
		return "", errors.New("exchange is of invalid type")
	}

	return exchangeString, nil
}

func getPair(c *gin.Context) (string, error) {
	pair, ok := c.Get(pairCtx)
	if !ok {
		return "", errors.New("pair not found")
	}

	pairString, ok := pair.(string)
	if !ok {
		return "", errors.New("pair is of invalid type")
	}

	return pairString, nil
}
