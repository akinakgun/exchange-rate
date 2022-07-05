package main

import (
	"exchange-rate/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	Server().Run(":5000")
}

func Server() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/:rateTo/:rateFrom/:amount/convertRate", handler.ConvertExchangeRatesHandler)
	r.GET("/getAll", handler.GetAllHandler)
	r.POST("/pullRate", handler.PullExchangeRateHandler)
	return r
}
