package handler

import (
	mongodb2 "exchange-rate/mongodb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var ConvertExchangeRatesHandler = func(c *gin.Context) {
	rateTo := c.Param("rateTo")
	rateFrom := c.Param("rateFrom")
	amount := c.Param("amount")

	result1, err := mongodb2.GetById(rateTo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	result2, err := mongodb2.GetById(rateFrom)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}

	rate := (result1.Value / result2.Value) * amountFloat

	c.JSON(http.StatusOK, rate)

}
