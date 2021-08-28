package ws

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed info.txt
var infoTXT string

/* TYPES DEFINITION */

type InfoWs interface {
	Get(context *gin.Context)
}

type DefaultInfoWs struct {
}

/* TYPES CONSTRUCTOR */

func NewDefaultInfoWs() *DefaultInfoWs {
	return &DefaultInfoWs{}
}

/* DefaultInfoWs METHODS */

func (ws DefaultInfoWs) Get(context *gin.Context) {
	context.Data(http.StatusOK, "text/text", []byte(infoTXT))
}
