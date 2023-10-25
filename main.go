package main

import (
	"example.com/lmwn-go-proj/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/covid/summary", handler.HandleCovidSummary)
	r.Run(":3000")
}
