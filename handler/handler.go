package handler

import (
	"net/http"

	"example.com/lmwn-go-proj/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.IService
}

func (h *Handler) HandleCovidSummary(c *gin.Context) {
	summary := h.Service.GetCovidDataSummary()
	c.JSON(http.StatusOK, summary)
}

var HandlerInstance = &Handler{
	Service: service.ServiceInstance,
}
