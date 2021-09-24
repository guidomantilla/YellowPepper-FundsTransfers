package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* TYPES DEFINITION */

type HealthWsResponse struct {
	Status string `json:"status"`
}

type HealthWs interface {
	Get(context *gin.Context)
}

type DefaultHealthWs struct {
}

/* TYPES CONSTRUCTOR */

func NewDefaultHealthWs() *DefaultHealthWs {
	return &DefaultHealthWs{}
}

/* DefaultMetricsWs METHODS */

func (ws *DefaultHealthWs) Get(context *gin.Context) {
	response := HealthWsResponse{
		Status: "UP",
	}
	context.JSON(http.StatusOK, response)
}
