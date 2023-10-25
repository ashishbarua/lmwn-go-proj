package handler

import (
	"net/http"

	"example.com/lmwn-go-proj/service"
	"github.com/gin-gonic/gin"
)

func HandleCovidSummary(c *gin.Context) {
	summary := service.GetCovidDataSummary()
	c.JSON(http.StatusOK, summary)
}
