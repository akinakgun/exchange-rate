package handler

import (
	mongodb2 "exchange-rate/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
)

var GetAllHandler = func(c *gin.Context) {

	result, err := mongodb2.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)

}
