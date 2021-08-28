package ws

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

/* TYPES DEFINITION */

type MetricsWs interface {
	Get(context *gin.Context)
}

type DefaultMetricsWs struct {
}

/* TYPES CONSTRUCTOR */

func NewDefaultMetricsWs() *DefaultMetricsWs {
	return &DefaultMetricsWs{}
}

/* DefaultMetricsWs METHODS */

func (ws DefaultMetricsWs) Get(context *gin.Context) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	context.JSON(http.StatusOK, mem)
}
