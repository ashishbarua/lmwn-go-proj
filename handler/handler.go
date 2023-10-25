package handler

import (
	"net/http"

	"example.com/lmwn-go-proj/service"
	"github.com/gin-gonic/gin"
)

var HandlerInstance = &Handler{
	Service: service.ServiceInstance,
}

type Handler struct {
	Service service.IService
}

func (h *Handler) HandleCovidSummary(c *gin.Context) {
	summary := h.Service.GetCovidDataSummary()
	c.JSON(http.StatusOK, summary)
}
